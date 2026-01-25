from textual.app import App, ComposeResult
from textual.widgets import Footer, Header, Static, Label, Checkbox, Collapsible, DataTable
from textual.containers import Vertical, Horizontal, Container, VerticalScroll
from helper.CategoryBox import FuncProgBar
import pandas as pd
import sys

class Pyfyle(App):

    CSS_PATH = "styles.tcss"
    BINDINGS = [("t", "table_togg", "Toggle Tables"),
                ("x", "time_togg", "Toggle between tottime,cumtime,ncalls")]
    TITLE = f"Pyfyle: Sorted by tottime"

    def __init__(self):
        super().__init__()
        self.mode = "tottime" # Initial mode
        self.tot_time = 0
        self.cum_time = 0


    def compose(self):
        csv_path = sys.argv[1]
        self.sub_title = f"Analyzing: {csv_path}"
        with Vertical(id="root"):
            yield Header()

            with VerticalScroll(id="content"):
                with Horizontal(id="func_cat"):
                    yield Checkbox("User-defined", id="cb_ud")
                    yield Checkbox("Builtin", id="cb_bi")
                    yield Checkbox("Frozen", id="cb_fr")
                    yield Checkbox("Others", id="cb_oth")

                
                df = pd.read_csv(csv_path)

                df.columns = df.columns.str.strip()

                df["ncalls"] = df["ncalls"].astype(str).apply(lambda x: int(x.split("/")[0]))

                builtin_mask = df['function'].str.contains('built-in method', na=False)
                builtins_df = df[builtin_mask].copy()

                c_ext_mask = df['function'].str.contains(r"method '", na=False)
                c_extensions_df = df[c_ext_mask].copy()

                frozen_mask = df["filename"].str.startswith("<frozen", na=False)
                frozen_df = df[frozen_mask].copy()

                user_func_df = df[~(builtin_mask | c_ext_mask | frozen_mask)].copy()

                tot_time = df['tottime'].sum()
                cum_time = df['cumtime'].sum()
                total_ncalls = df['ncalls'].sum()
                self.tot_time = tot_time
                self.cum_time = cum_time
                self.total_ncalls = total_ncalls

                yield FuncProgBar("User-defined", user_func_df, _id="panel1", tot_time=tot_time, cum_time=cum_time, total_ncalls=total_ncalls)
                yield FuncProgBar("Builtin", builtins_df, _id="panel2", tot_time=tot_time, cum_time=cum_time, total_ncalls=total_ncalls)
                yield FuncProgBar("Frozen", frozen_df, _id="panel3", tot_time=tot_time, cum_time=cum_time, total_ncalls=total_ncalls)
                yield FuncProgBar("Others", c_extensions_df, _id="panel4", tot_time=tot_time, cum_time=cum_time, total_ncalls=total_ncalls)

            yield Footer()


    def action_table_togg(self) -> None:                  
        for collapsible in self.query(Collapsible):
            collapsible.collapsed = not collapsible.collapsed

    def action_time_togg(self) -> None:
        modes = ["tottime", "cumtime", "ncalls"]
        
        current_index = modes.index(self.mode)
        self.mode = modes[(current_index + 1) % len(modes)]

        if self.mode == "ncalls":
            new_total = self.total_ncalls 
        elif self.mode == "cumtime":
            new_total = self.cum_time
        else:
            new_total = self.tot_time

        for category in self.query(FuncProgBar):
            category.rebuild_bars(self.mode)

        self.title = f"Pyfyle: Sorted by {self.mode}"

    def on_checkbox_changed(self, event) -> None:
        if event.checkbox.id == "cb_ud":
            panel = self.query_one(".panel1")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"
        if event.checkbox.id == "cb_bi":
            panel = self.query_one(".panel2")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"
        if event.checkbox.id == "cb_fr":
            panel = self.query_one(".panel3")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"
        if event.checkbox.id == "cb_oth":
            panel = self.query_one(".panel4")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"

    def on_mount(self) -> None:
        self.theme = "nord"




if __name__ == "__main__":
    app = Pyfyle()
    app.run()