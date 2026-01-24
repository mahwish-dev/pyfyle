from textual.widgets import Collapsible, Label, DataTable, ProgressBar
from textual.widget import Widget

import pandas as pd

class FuncProgBar(Widget):
	def __init__(self, title:str, df:pd.DataFrame, _id:str, tot_time:int, cum_time:int):
		super().__init__()
		self.title = title
		self.df = df
		self._id = _id
		self.tot_time = tot_time
		self.cum_time = cum_time

	def compose(self):
		if self.df.empty:
			with Collapsible(title=self.title, collapsed=True, id=self._id):
				yield Label(f"No {self.title} functions.")
		else: 

			for row in self.df.itertuples():

				prog_bar = ProgressBar(total=self.tot_time, classes="progress-bar")
				prog_bar.update(progress=row.tottime)
				prog_bar.tottime_val = row.tottime
				prog_bar.cumtime_val = row.cumtime
				yield prog_bar

			with Collapsible(title=self.title, collapsed=True, id=self._id):
				table = DataTable(classes="tables")
				table.add_columns(*self.df.columns.astype(str))
				table.add_rows(
					self.df.astype(str).to_numpy().tolist()
				)

		
				yield table


