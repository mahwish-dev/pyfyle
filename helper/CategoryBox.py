from textual.widgets import Collapsible, Label, DataTable, ProgressBar
from textual.widget import Widget
from textual.containers import Grid, Horizontal

import pandas as pd

class FuncProgBar(Widget):
	def __init__(self, title:str, df:pd.DataFrame, _id:str, tot_time:int, cum_time:int):
		super().__init__()
		self.title = title
		self.df = df
		self._id = _id
		self.tot_time = tot_time
		self.cum_time = cum_time
		self.mode = "tottime"

	def prettify_text(self, txt:str):
		if self.title == "Builtin":
			return txt[16:-1]
		elif self.title == "Others":
			return txt[7:-8]
		else:
			return txt

	def compose(self):
		if self.df.empty:
			with Collapsible(title=self.title, collapsed=True, id=self._id):
				yield Label(f"No {self.title} functions.")
		else: 

			self.df = self.df.sort_values(by=self.mode, ascending=False)

			for row in self.df.itertuples():
				with Horizontal():
					prog_bar = ProgressBar(total=self.tot_time, classes="progress-bar", show_eta=False)
					prog_bar.update(progress=row.tottime)
					prog_bar.tottime_val = row.tottime
					prog_bar.cumtime_val = row.cumtime
					prog_bar.ncalls_val = row.ncalls
					yield prog_bar
					yield Label(f" {self.prettify_text(row.function)}")


			with Collapsible(title=f"[b]{self.title}[/b]   Total: [i]{self.df['tottime'].sum():.3f}[/i]", collapsed=True, classes=self._id):
				table = DataTable(classes="tables")
				table.add_columns(*self.df.columns.astype(str))
				table.add_rows(
					self.df.astype(str).to_numpy().tolist()
				)

		
				yield table


	def rebuild_bars(self, mode) -> None:
		self.mode = mode

		self.df = self.df.sort_values(by=self.mode, ascending=False)

		container = self.query_one(f'.{self._id}')

		if self.mode == "ncalls":
			current_total = self.df['ncalls'].sum()
		elif self.mode == "cumtime":
			current_total = self.df['cumtime'].sum()
		else:
			current_total = self.df['tottime'].sum()

		self.remove_children()

		with self.app.batch_update():

			for row in self.df.itertuples():

				horiz = Horizontal()
				self.mount(horiz)

				prog_bar = ProgressBar(total=self.tot_time, classes="progress-bar", show_eta=False)
				prog_bar.update(progress=row.tottime)
				prog_bar.tottime_val = row.tottime
				prog_bar.cumtime_val = row.cumtime
				prog_bar.ncalls_val = row.ncalls
						
				horiz.mount(prog_bar)
				horiz.mount(Label(f" {self.prettify_text(row.function)}"))

				

			collap = Collapsible(title=f"[b]{self.title}[/b] | Total: [i]{self.df['tottime'].sum():.3f}[/i]", collapsed=True, classes=self._id)

			table = DataTable(classes="tables")
			table.add_columns(*self.df.columns.astype(str))
			table.add_rows(
				self.df.astype(str).to_numpy().tolist()
			)
			
			self.mount(collap)
			collap.mount(table)

