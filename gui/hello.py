import tkinter 

#tryna see if I can just make a fixed size window
fixed_width = 1000
fixed_height = 800 

root = tkinter.Tk() 

manga_var = tkinter.StringVar()

def submitAction():
    mangaName = manga_var.get()
    print(f"the name of the manga is: {mangaName}")
    manga_var.set("")
    
def getRecommendation():
    mangaName = manga_var.get()
    #go call the thing here
    #result = whatever the response is here
    return "the recommendation here"

def getSynopsis():
    mangaName = manga_var.get()
    #go call the thing here
    #synopsis = whatever the response is here
    return "the synopsis"

#geometry is used for setting the window sized automatically
root.geometry(f"{fixed_width}x{fixed_height}")
#to make it resizable or not 
root.resizable(False,False)
w = tkinter.Label(root, text="lebron james")

#the basic entry stuff
manga_name =  tkinter.Label(root, text="Name of the Manga:")
manga_entry = tkinter.Entry(root, textvariable=manga_var)

#the buttons of what I can do 
get_rec_button = tkinter.Button(root, text="Get Recommendation", command=getRecommendation)
get_synopsis = tkinter.Button(root, text="Get Synopsis", command=getSynopsis)

#the grid stuff, basically how its set up
manga_name.grid(row=0,column=0)
manga_entry.grid(row=0,column=1)
get_rec_button.grid(row=1,column=0)
get_synopsis.grid(row=1,column=1)

root.mainloop()

#see if this only updates in skib branch