from textual.app import App, ComposeResult
from textual.widgets import Footer, Header, Static, Label, Checkbox, Collapsible, DataTable
from textual.containers import Vertical, Horizontal, Container, VerticalScroll
from helper.CategoryBox import FuncProgBar
import pandas as pd

class Pyfyle(App):

    CSS_PATH = "styles.tcss"
    BINDINGS = [("b", "builtin_togg", "Toggle builtin functions")]
    TITLE = "Pyfyle"

    def __init__(self):
        super().__init__()


    def compose(self):
        with Vertical(id="root"):
            yield Header()

            with VerticalScroll(id="content"):
                yield Static("Function categories:")
                with Horizontal(id="func_cat"):
                    yield Checkbox("User-defined", id="cb_ud")
                    yield Checkbox("Builtin", id="cb_bi")
                    yield Checkbox("Others", id="cb_oth")

                df = pd.read_csv("profile_results.csv")

                yield FuncProgBar("User-defined", df, _id="panel1")
                yield FuncProgBar("Builtin", df, _id="panel2")
                yield FuncProgBar("Others", df, _id="panel3")

            yield Footer()


    def action_builtin_togg(self) -> None:      
        pass 
    def action_internals_togg(self) -> None:
        pass
    def on_checkbox_changed(self, event) -> None:
        if event.checkbox.id == "cb_ud":
            panel = self.query_one("#panel1")
            panel.display = "none" if panel.styles.display == "block" else "block"
        if event.checkbox.id == "cb_bi":
            panel = self.query_one("#panel2")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"
        if event.checkbox.id == "cb_oth":
            panel = self.query_one("#panel3")
            panel.styles.display = "none" if panel.styles.display == "block" else "block"




if __name__ == "__main__":
    app = Pyfyle()
    app.run()
