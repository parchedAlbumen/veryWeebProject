import tkinter 
import requests
import json
from PIL import Image, ImageTk
from io import BytesIO

#tryna see if I can just make a fixed size window
fixed_width = 1000
fixed_height = 800 
init_msg = "type a manga,,, press one of the buttons,,, see what you get!"

root = tkinter.Tk() 
manga_var = tkinter.StringVar()

root.title("Manga wowzers!")

#To submit action
def submitAction():
    mangaName = manga_var.get()
    print(f"mangaName: {mangaName}")
    manga_var.set("") #hopefully self-explanatory 
    
#To get a synopsis of the manga, I am going to add an LLM to get lebron james to summarize this in his own
def getSynopsis():
    dataName = {"mangaName": manga_var.get()}

    response = requests.post("http://localhost:3333/skibidiRizzlerSigmaMale", json=dataName) #get response
    data = response.json()

    msg.config(text=data["response"]) #cuz json thingz
    updateImage(data["imageurl"])
    manga_var.set("")

#To get score, I am going to get an LLM to summarize this, so it would seem cooler!
def getScore():
    dataName = {"mangaName": manga_var.get()}
    response = requests.post("http://localhost:3333/getScore", json=dataName)
    data = response.json()

    msg.config(text=data["response"])
    updateImage(data["imageurl"])
    manga_var.set("")

#To get a recommendation from the given thing
def getRecommendation():
    dataName = {"mangaName": manga_var.get()}
    response = requests.post("http://localhost:3333/getRec", json=dataName)
    data = response.json()

    manga_var.set("")

#Updates the image per search
def updateImage(url):
    print(url)
    if len(url) > 0: 
        resp = requests.get(url)
        image = Image.open(BytesIO(resp.content))
        photo = ImageTk.PhotoImage(image)
        manga_photo.config(image=photo)  #shows the image
        manga_photo.image = photo   #keeps the image alive in the manga_photo attribute, to avoid garbage collector
    else: 
        print("no images i guess")


#geometry is used for setting the window sized automatically
root.geometry(f"{fixed_width}x{fixed_height}")
#to make it resizable or not, here we chose not to lol
root.resizable(False,False)
w = tkinter.Label(root, text="lebron james")

#the basic entry stuff
manga_name =  tkinter.Label(root, text="Name of the Manga:")
manga_entry = tkinter.Entry(root, textvariable=manga_var)

#the buttons of what I can do 
get_score= tkinter.Button(root, text="Get Score", command=getRecommendation)
get_synopsis = tkinter.Button(root, text="Get Synopsis", command=getSynopsis)

#make the message thing here
msg = tkinter.Label(root, text=init_msg, wraplength=300) 
manga_photo = tkinter.Label(root, image="")

#the grid stuff, basically how its set up
manga_name.grid(row=0,column=0)
manga_entry.grid(row=0,column=1)
get_score.grid(row=1,column=0)
get_synopsis.grid(row=1,column=1)
msg.grid(row=0, column=3,padx=100)
manga_photo.grid(row=3, column=0)

root.mainloop()
