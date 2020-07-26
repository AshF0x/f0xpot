import time
import board
import neopixel
import sys

# Code is working for RGB, i didn't have RGBW to try, but it should not be hard to modify

# if you want to have specific pattern for specific Country/Continent uncomment this
# eg. if attack comes from china (CN) you would access it like: python3 pixelAlert.py CN
# !!! country codes must be UPPERCASE, or lowercase if you write them like that, but since the API returns capital letters... youget the idea
# For this to work ypu need to uncomment some code at the end of the file

# country = str(sys.argv[1]);

pixel_pin = board.D18       #remember to set the pin number
num_pixels = 16             #number of LEDs
ORDER = neopixel.GRBW
pixels = neopixel.NeoPixel(
    pixel_pin, num_pixels, brightness=0.5, auto_write=False, pixel_order=ORDER
)

def police(rpt):                     # rpt <- integer that tells how many times to repeat
    for i in range(rpt):
        for a in range(4) :          # repeat this pattern (X) times for 8 LEDs
            pixels.fill((0,0,0,0))     # making sure LEDs are turned off
            pixels.show()            # telling LEDs to turn off
            pixels[0]=((255,0,0,0))   #should be red
            pixels[1]=((255,0,0,0))    #should be red
            pixels[2]=((255,0,0,0))    #should be red
            pixels[3]=((255,0,0,0))    #should be red
            pixels[4]=((255,0,0,0))    #should be red
            pixels[5]=((255,0,0,0))    #should be red
            pixels[6]=((255,0,0,0))    #should be red
            pixels[7]=((255,0,0,0))    #should be red
            pixels.show()
            time.sleep(0.02)
            pixels.fill((0,0,0,0))
            pixels.show()
        for b in range(4):           # repeat this pattern (X) times for the other 8 LEDs
            pixels.fill((0,0,0,0))
            pixels.show()
            pixels[8]=((0,255,0,0))    #should be Blue
            pixels[9]=((0,255,0,0))     #should be Blue
            pixels[10]=((0,255,0,0))     #should be Blue
            pixels[11]=((0,255,0,0))     #should be Blue
            pixels[12]=((0,255,0,0))     #should be Blue
            pixels[13]=((0,255,0,0))     #should be Blue
            pixels[14]=((0,255,0,0))     #should be Blue
            pixels[15]=((0,255,0,0))     #should be Blue
            pixels.show()
            time.sleep(0.02)
            pixels.fill((0,0,0))
            pixels.show()
        time.sleep(.2)
        for c in range(3) :
            pixels.fill((0,0,0,0))
            pixels.show()
            pixels.fill((0,0,0,255))    #should be semething like white for White
            pixels.show()
            time.sleep(.2)
    pixels.fill((0,0,0))
    pixels.show()

police(3)       #repeat  times, if you want to get repeat times from imput you just need to write repeat instead of this 3

# Uncoment this block and, coment the line up it you want to use it with specicif country 
"""
if country == "CN" :
    pixels.fill((0,0,0,0))     #should be Blue
    pixels.show()
    time.sleep(0.2)
    pixels.fill((0,255,0,0))     #should be Blue
    pixels.show()
    time.sleep(0.5)
    pixels.fill((0,0,0,0))
    pixels.show()
    time.sleep(0.2)
    pixels.fill((0,0,255,0))     #should be Blue
    pixels.show()
    time.sleep(0.5)
    pixels.fill((0,0,0,0))
    pixels.show()
    time.sleep(0.2)
elif country == 'DE': 
    pixels.fill((11, 123, 23,254))     #should be some color
    pixels.show()
    time.sleep(0.02)
    pixels.fill((0,0,0,0))
    pixels.fill((11, 123, 23,55))     #should be some color
    pixels.show()
    pixels.fill((0,0,0,0))
    pixels.show()
    time.sleep(0.02)
    pixels.fill((11, 123, 23,0))     #should be some color
    pixels.show()
    time.sleep(0.02)
    pixels.fill((0,0,0,0))
    pixels.show()
    time.sleep(0.02)
    pixels.fill((124,25,51,121))     #should be some color
    pixels.show()
    time.sleep(0.5)
    pixels.fill((0,0,0,0))
    pixels.show()
else :
    police(3)
"""
