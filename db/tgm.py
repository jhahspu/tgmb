import tkinter as tk
from tkinter import StringVar, ttk
from tkinter import messagebox, filedialog
import csv, sqlite3


conn = sqlite3.connect("tgm.db")
# TEST IN MEMORY (uncommenct next line)
# conn = sqlite3.connect(":memory:")
c = conn.cursor()
CREATE_TBL = """
CREATE TABLE IF NOT EXISTS mvs (
  tmdb int PRIMARY KEY,
  title varchar(255),
  tagline varchar(255),
  release string,
  runtime int,
  genres varchar(255),
  overview text,
  poster varchar(255),
  backdrop varchar(255),
  trailers varchar(255)
)
"""
c.execute(CREATE_TBL)
conn.commit()

def clear_inputs():
  e1.set('')
  e2.set('')
  e3.set('')
  e4.set('')
  e5.set('')
  e6.set('')
  e7.set('')
  e8.set('')
  e9.set('')
  e10.set('')
  q.set('')


def selected_row(event):
  # print("[double clicked a row] ")
  # rowid = trv.identify_row(event.y)
  item = trv.item(trv.focus())
  # item['values'][0]
  e1.set(item['values'][0])
  e2.set(item['values'][1])
  e3.set(item['values'][2])
  e4.set(item['values'][3])
  e5.set(item['values'][4])
  e6.set(item['values'][5])
  e7.set(item['values'][6])
  e8.set(item['values'][7])
  e9.set(item['values'][8])
  e10.set(item['values'][9])

def search_all():
  get_all(q.get())
  clear_inputs()

def get_all(app=""):
  if app != "":
    c.execute("""
    SELECT * FROM mvs WHERE app=:app
    """, {'app': app})
  else:
    c.execute("""
    SELECT * FROM mvs
    """
    )
  rows = c.fetchall()
  trv.delete(*trv.get_children())
  for i in rows:
    trv.insert('', 'end', values=i)
  clear_inputs()

def app_add():
  with conn:
    c.execute("""
    INSERT INTO mvs
      VALUES (:tmdb, :title, :tagline, :release, :runtime, :genres, :overview, :poster, :backdrop, :trailers)
    """, {
      'tmdb': e1.get(),
      'title': e2.get(),
      'tagline': e3.get(),
      'release': e4.get(),
      'runtime': e5.get(),
      'genres': e6.get(),
      'overview': e7.get(),
      'poster': e8.get(),
      'backdrop': e9.get(),
      'trailers': e10.get()
      }
    )
  get_all()

def app_update():
  with conn:
    c.execute("""
    UPDATE mvs
      SET title=:title, tagline=:tagline, release=:release, runtime=:runtime, genres=:genres, overview=:overview, poster=:poster, backdrop=:backdrop, trailers=:trailers
      WHERE tmdb=:tmdb
    """, {
        'tmdb': e1.get(),
        'title': e2.get(),
        'tagline': e3.get(),
        'release': e4.get(),
        'runtime': e5.get(),
        'genres': e6.get(),
        'overview': e7.get(),
        'poster': e8.get(),
        'backdrop': e9.get(),
        'trailers': e10.get(),
        }
    )
  get_all()

def app_delete():
  if messagebox.askyesno("Confirmation needed", "Are you sure you to delete App?"):
    with conn:
      c.execute("""
      DELETE FROM mvs
        WHERE tmdb=:tmdb
      """, {'tmdb': e1.get()}
      )
    get_all()

def insertFromTSV():
  with open('movies.tsv', 'r', encoding="utf8") as f:
    dr = csv.DictReader(f, delimiter="\t")
    to_db = [(i['tmdb'], i['title'], i['tagline'], i['release'], i['runtime'], i['genres'], i['overview'], i['poster'], i['backdrop'], i['trailers']) for i in dr]
  with conn:
    c.execute("DELETE FROM mvs")
    c.executemany("INSERT INTO mvs (tmdb, title, tagline, release, runtime, genres, overview, poster, backdrop, trailers) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", to_db)
  get_all()

root = tk.Tk()

wrapper1 = tk.LabelFrame(root, text="List")
wrapper2 = tk.LabelFrame(root, text="Search")
wrapper3 = tk.LabelFrame(root, text="Update")

wrapper1.pack(fill="both", expand="yes", padx=20, pady=10)
wrapper2.pack(fill="both", expand="yes", padx=20, pady=10)
wrapper3.pack(fill="both", expand="yes", padx=20, pady=10)

# List: Wrapper 1
trv = ttk.Treeview(wrapper1, columns=(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), show="headings", height="6")
trv.heading(1, text="tMDb")
trv.heading(2, text="Title")
trv.heading(3, text="Tagline")
trv.heading(4, text="Release")
trv.heading(5, text="Runtime")
trv.heading(6, text="Genres")
trv.heading(7, text="Overview")
trv.heading(8, text="Poster")
trv.heading(9, text="Backdrop")
trv.heading(10, text="Trailers")
trv.pack(fill="both")

trv.bind('<Double 1>', selected_row)

# Search: Wrapper 2
lbl = tk.Label(wrapper2, text="Search")
lbl.pack(side=tk.LEFT, padx=10)

q = StringVar()
ent = tk.Entry(wrapper2, textvariable=q)
ent.pack(side=tk.LEFT, padx=6)

btn = tk.Button(wrapper2, text="Search", command=search_all)
btn.pack(side=tk.LEFT, padx=6)

# Update: Wrapper 3
e1 = StringVar()
lbl1 = tk.Label(wrapper3, text="tMDb")
lbl1.grid(row=0, column=0, padx=5, pady=3)
ent1 = tk.Entry(wrapper3, textvariable=e1)
ent1.grid(row=0, column=1, padx=5, pady=3)


e2 = StringVar()
lbl2 = tk.Label(wrapper3, text="Title")
lbl2.grid(row=1, column=0, padx=5, pady=3)
ent2 = tk.Entry(wrapper3, textvariable=e2)
ent2.grid(row=1, column=1, padx=5, pady=3)


e3 = StringVar()
lbl3 = tk.Label(wrapper3, text="Tagline")
lbl3.grid(row=2, column=0, padx=5, pady=3)
ent3 = tk.Entry(wrapper3, textvariable=e3)
ent3.grid(row=2, column=1, padx=5, pady=3)


e4 = StringVar()
lbl4 = tk.Label(wrapper3, text="Release")
lbl4.grid(row=3, column=0, padx=5, pady=3)
ent4 = tk.Entry(wrapper3, textvariable=e4)
ent4.grid(row=3, column=1, padx=5, pady=3)


e5 = StringVar()
lbl5 = tk.Label(wrapper3, text="Runtime")
lbl5.grid(row=4, column=0, padx=5, pady=3)
ent5 = tk.Entry(wrapper3, textvariable=e5)
ent5.grid(row=4, column=1, padx=5, pady=3)


e6 = StringVar()
lbl6 = tk.Label(wrapper3, text="Genres")
lbl6.grid(row=5, column=0, padx=5, pady=3)
ent6 = tk.Entry(wrapper3, textvariable=e6)
ent6.grid(row=5, column=1, padx=5, pady=3)


e7 = StringVar()
lbl7 = tk.Label(wrapper3, text="Overview")
lbl7.grid(row=6, column=0, padx=5, pady=3)
ent7 = tk.Entry(wrapper3, textvariable=e7)
ent7.grid(row=6, column=1, padx=5, pady=3)


e8 = StringVar()
lbl8 = tk.Label(wrapper3, text="Poster")
lbl8.grid(row=7, column=0, padx=5, pady=3)
ent8 = tk.Entry(wrapper3, textvariable=e8)
ent8.grid(row=7, column=1, padx=5, pady=3)


e9 = StringVar()
lbl9 = tk.Label(wrapper3, text="Backdrop")
lbl9.grid(row=8, column=0, padx=5, pady=3)
ent9 = tk.Entry(wrapper3, textvariable=e9)
ent9.grid(row=8, column=1, padx=5, pady=3)


e10 = StringVar()
lbl10 = tk.Label(wrapper3, text="Trailers")
lbl10.grid(row=9, column=0, padx=5, pady=3)
ent10 = tk.Entry(wrapper3, textvariable=e10)
ent10.grid(row=9, column=1, padx=5, pady=3)



addBtn = tk.Button(wrapper3, text="Create", command=app_add)
addBtn.grid(row=11, column=1, padx=5, pady=3)

updateBtn = tk.Button(wrapper3, text="Update", command=app_update)
updateBtn.grid(row=11, column=2, padx=5, pady=3)

deleteBtn = tk.Button(wrapper3, text="Detele", command=app_delete)
deleteBtn.grid(row=11, column=3, padx=5, pady=3)

getFromTSVBtn = tk.Button(wrapper3, text="TSV Import", command=insertFromTSV)
getFromTSVBtn.grid(row=11, column=5, padx=5, pady=3)

get_all()

root.title("Top Good Movies")
root.geometry("800x700")
root.mainloop()

conn.close()