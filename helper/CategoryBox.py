from textual.widgets import Collapsible, Label, DataTable, ProgressBar
from textual.widget import Widget
from textual.containers import Grid, Horizontal, Vertical

import pandas as pd

class FuncProgBar(Widget):
	def __init__(self, title:str, df:pd.DataFrame, _id:str, tot_time:int, cum_time:int, total_ncalls:int):
		super().__init__()
		self.title = title
		self.df = df
		self._id = _id
		self.tot_time = tot_time
		self.cum_time = cum_time
		self.mode = "tottime"
		self.total_ncalls = total_ncalls

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

			i = 0
			with Vertical(classes="bars-area"):
				for row in self.df.itertuples():
					with Horizontal(classes="bars-area"):
						prog_bar = ProgressBar(total=self.tot_time, classes="progress-bar", show_eta=False)
						prog_bar.update(progress=row.tottime)
						prog_bar.tottime_val = row.tottime
						prog_bar.cumtime_val = row.cumtime
						prog_bar.ncalls_val = row.ncalls
						yield prog_bar
						yield Label(f" {self.prettify_text(row.function)}")
						i += 1
					if i == 5:
						break


			with Collapsible(title=f"[b]{self.title}[/b] | Total: [i]{self.df['tottime'].sum():.3f}[/i]", collapsed=True, classes=self._id):
				table = DataTable(classes="tables")
				table.add_columns(*self.df.columns.astype(str))
				table.add_rows(
					self.df.astype(str).to_numpy().tolist()
				)

		
				yield table

	def build_prog_bar(self, row) -> ProgressBar:
		# "tottime", "cumtime", "ncalls"
		if self.mode == "tottime":
			prog_bar = ProgressBar(total=self.tot_time, classes="progress-bar", show_eta=False)
			prog_bar.update(progress=row.tottime)
		elif self.mode == "cumtime":
			prog_bar = ProgressBar(total=self.cum_time, classes="progress-bar", show_eta=False)
			prog_bar.update(progress=row.cumtime)
		else:
			prog_bar = ProgressBar(total=self.total_ncalls, classes="progress-bar", show_eta=False)
			prog_bar.update(progress=row.ncalls)

		prog_bar.tottime_val = row.tottime
		prog_bar.cumtime_val = row.cumtime
		prog_bar.ncalls_val = row.ncalls

		return prog_bar

	def rebuild_bars(self, mode) -> None:
		self.mode = mode
		self.df = self.df.sort_values(by=self.mode, ascending=False)

		was_collapsed = self.query_one(Collapsible).collapsed

		bars_area = self.query_one(".bars-area", Vertical)
		bars_area.remove_children()

		current_total = self.df[self.mode].sum()

		i = 0
		with self.app.batch_update():
			for row in self.df.itertuples():

				horiz = Horizontal()
				bars_area.mount(horiz)

				prog_bar = self.build_prog_bar(row)
						
				horiz.mount(prog_bar)
				horiz.mount(Label(f" {self.prettify_text(row.function)}"))

				i += 1
				if i == 5:
					break

		table = self.query_one(".tables", DataTable)
		table.clear(columns=True) # Clear columns and rows
		table.add_columns(*self.df.columns.astype(str))
		table.add_rows(self.df.astype(str).to_numpy().tolist())

		collap = self.query_one(Collapsible)
		collap.title = f"[b]{self.title}[/b] | Total: [i]{self.df[self.mode].sum():.3f}[/i]"