import tkinter 

root = tkinter.Tk() 

#tryna see if I can just make a fixed size window
fixed_width = 1000
fixed_height = 800 

#geometry is used for setting the window sized automatically
root.geometry(f"{fixed_width}x{fixed_height}")
#to make it resizable or not 
root.resizable(False,False)
w = tkinter.Label(root, text="lebron james")
w.pack() 
root.mainloop()

