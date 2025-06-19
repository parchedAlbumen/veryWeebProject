import tkinter as tk #literally just imports and renames it as tk
from tkinter import ttk #themed tkinter, modern looking widget thingy
from tkinter import font #style font, self explanatory
import gui_commands as gc

#window
root = tk.Tk() 
root.title("Manga wowzers!") 
# root.configure(bg="#353839") #use onyx colour lol

#set ups
#tryna see if I can just make a fixed size window
fixed_width = 800
fixed_height = 600 
root.geometry(f"{fixed_width}x{fixed_height}")
#to make it resizable or not, here we chose not to lol
root.resizable(False,False)
 
#theme styling kinda
style = ttk.Style()
style.theme_use('clam') #i can't tell if this is changing anything honestly

# # Configure custom styles  ###note: figure these out 
# style.configure('Title.TLabel',
#                 background='#2c3e50',
#                 foreground='#ecf0f1',
#                 font=('Segoe UI', 24, 'bold'))

# style.configure('Header.TLabel',
#                 background='#2c3e50',
#                 foreground='#3498db',
#                 font=('Segoe UI', 12, 'bold'))

# style.configure('Info.TLabel',
#                 background='#34495e',
#                 foreground='#ecf0f1',
#                 font=('Segoe UI', 10),
#                 relief='flat',
#                 padding=10)

# style.configure('Modern.TButton',
#                 font=('Segoe UI', 10, 'bold'),
#                 padding=10)

# style.map('Modern.TButton',
#           background=[('active', '#3498db'),
#                      ('pressed', '#2980b9')])

# style.configure('Search.TEntry',
#                 font=('Segoe UI', 11),
#                 padding=5)

# Main frame, similar to a div, its like a container
main_frame = tk.Frame(root, bg="#353839", padx=30, pady=20)
main_frame.pack(fill='both', expand=True)

# Title, we want to put it inside the main frame 
title_label = ttk.Label(main_frame, text="Manga Stuff", style='Title.TLabel')
title_label.pack(pady=(0, 30)) #so that it shows something like a shadow

# Search section frame
search_frame = tk.Frame(main_frame, bg='#34495e', relief='raised', bd=2)
search_frame.pack(fill='x', pady=(0, 20), padx=20)

# Search section content
search_content = tk.Frame(search_frame, bg='#34495e', padx=20, pady=15)
search_content.pack(fill='x')

manga_name_label = ttk.Label(search_content, text="Enter Manga Name:", style='Header.TLabel')
manga_name_label.pack(anchor='w', pady=(0, 5))

# Entry with StringVar (you'll need to define manga_var)
manga_var = tk.StringVar() #the actual text
manga_entry = ttk.Entry(search_content, textvariable=manga_var, style='Search.TEntry', width=50) #the text box
manga_entry.pack(fill='x', pady=(0, 15))

# Buttons frame
buttons_frame = tk.Frame(search_content, bg='#34495e')
buttons_frame.pack(fill='x')

# Create buttons with modern styling
get_score = ttk.Button(buttons_frame, text="📊 Get Score", style='Modern.TButton')
get_score.pack(side='left', padx=(0, 10))

get_synopsis = ttk.Button(buttons_frame, text="📖 Get Synopsis", style='Modern.TButton')
get_synopsis.pack(side='left', padx=(0, 10))

get_rec = ttk.Button(buttons_frame, text="💡 Get Recommendation", style='Modern.TButton')
get_rec.pack(side='left')

# Results section
results_frame = tk.Frame(main_frame, bg='#2c3e50')
results_frame.pack(fill='both', expand=True, pady=(20, 0))

# Create two columns for results
left_column = tk.Frame(results_frame, bg='#2c3e50')
left_column.pack(side='left', fill='both', expand=True, padx=(0, 10))

right_column = tk.Frame(results_frame, bg='#2c3e50')
right_column.pack(side='right', fill='y', padx=(10, 0))

# Message display with scrollbar
msg_frame = tk.Frame(left_column, bg='#34495e', relief='sunken', bd=2)
msg_frame.pack(fill='both', expand=True)

msg_label = tk.Label(msg_frame, 
                    text="Welcome to Manga Explorer!\n\nEnter a manga name above and click one of the buttons to get started.",
                    bg='#34495e',
                    fg='#ecf0f1',
                    font=('Segoe UI', 11),
                    wraplength=400,
                    justify='left',
                    padx=10,
                    pady=10)
msg_label.pack(fill='both', expand=True)

# Image display area
image_frame = tk.Frame(right_column, bg='#34495e', relief='sunken', bd=2, width=200, height=300)
image_frame.pack(fill='y', expand=True)
image_frame.pack_propagate(False)  # Maintain fixed size

manga_photo = tk.Label(image_frame, 
                      text="📱\nManga Cover\nWill Appear Here",
                      bg='#34495e',
                      fg='#95a5a6',
                      font=('Segoe UI', 10))
manga_photo.pack(expand=True)

#wrapper commands:
def rec_wrapper():
    gc.getRecommendation(msg_label, manga_photo, manga_var)

def score_wrapper():
    gc.getScore(msg_label, manga_photo, manga_var)  

def synopsis_wrapper():
    gc.getSynopsis(msg_label, manga_photo, manga_var, root)  

get_rec.configure(command=rec_wrapper)
get_score.configure(command=score_wrapper)
get_synopsis.configure(command=synopsis_wrapper)

root.mainloop()