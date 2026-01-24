from textual.widgets import Collapsible, Label, DataTable, ProgressBar
from textual.widget import Widget

import pandas as pd

class FuncProgBar(Widget):
	def __init__(self, title:str, df:pd.DataFrame, _id:str):
		super().__init__()
		self.title = title
		self.df = df
		self._id = _id

	def compose(self):

		for _ in self.df['function']:
			yield ProgressBar(total=100)

		with Collapsible(title=self.title, collapsed=True, id=self._id):
			table = DataTable()
			table.add_columns(*self.df.columns.astype(str))
			table.add_rows(
				self.df.astype(str).to_numpy().tolist()
			)

		if self.df.empty:
			yield Label(f"No {self.title} functions.")
		else: 
			yield table

