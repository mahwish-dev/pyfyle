from textual.app import App, ComposeResult
from textual.widgets import Footer, Header, Static, Label, Checkbox, Collapsible, DataTable
from textual.containers import Vertical, Horizontal, Container, VerticalScroll
from helper.CategoryBox import FuncProgBar
import pandas as pd
import sys

class Pyfyle(App):

    CSS_PATH = "styles.tcss"
    BINDINGS = [("t", "table_togg", "Toggle Tables"),
                ("x", "time_togg", "Toggle between tottime and cumtime")]
    TITLE = "Pyfyle"

    def __init__(self):
        super().__init__()
        self.mode = "tottime" # Initial mode
        self.tot_time = 0
        self.cum_time = 0


    def compose(self):
        with Vertical(id="root"):
            yield Header()

            with VerticalScroll(id="content"):
                yield Static("Function categories:")
                with Horizontal(id="func_cat"):
                    yield Checkbox("User-defined", id="cb_ud")
                    yield Checkbox("Builtin", id="cb_bi")
                    yield Checkbox("Others", id="cb_oth")

                csv_path = sys.argv[1]
                df = pd.read_csv(csv_path)

                df.columns = df.columns.str.strip()

                builtin_mask = df['function'].str.contains('built-in method', na=False)
                builtins_df = df[builtin_mask].copy()

                c_ext_mask = df['function'].str.contains(r"<method '", na=False)
                c_extensions_df = df[c_ext_mask].copy()

                user_func_df = df[~(builtin_mask | c_ext_mask)].copy()

                tot_time = df['tottime'].sum()
                cum_time = df['cumtime'].sum()
                self.tot_time = tot_time
                self.cum_time = cum_time

                yield FuncProgBar("User-defined", user_func_df, _id="panel1", tot_time=tot_time, cum_time=cum_time)
                yield FuncProgBar("Builtin", builtins_df, _id="panel2", tot_time=tot_time, cum_time=cum_time)
                yield FuncProgBar("Others", c_extensions_df, _id="panel3", tot_time=tot_time, cum_time=cum_time)

            yield Footer()


    def action_table_togg(self) -> None:                  
        for collapsible in self.query(Collapsible):
            collapsible.collapsed = not collapsible.collapsed

    def action_time_togg(self) -> None:
        # 1. Flip the mode
        self.mode = "cumtime" if self.mode == "tottime" else "tottime"

        # 2. Update the bars using the data we "tucked away" inside them
        for bar in self.query(".progress-bar"):
            if self.mode == "cumtime":
                bar.total = self.cum_time
                bar.update(progress=bar.cumtime_val)
            else:
                bar.total = self.tot_time
                bar.update(progress=bar.tottime_val)

    def on_checkbox_changed(self, event) -> None:
        if event.checkbox.id == "cb_ud":
            panel = self.query_one("#panel1")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"
        if event.checkbox.id == "cb_bi":
            panel = self.query_one("#panel2")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"
        if event.checkbox.id == "cb_oth":
            panel = self.query_one("#panel3")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"




if __name__ == "__main__":
    app = Pyfyle()
    app.run()