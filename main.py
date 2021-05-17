#!/usr/bin/python3

# pip3 install pycrypto

from Crypto.Cipher import AES
import re
from sys import argv
import urllib.request

wsdir = 'workspace/'

EncKeyProvider = ""
EncKey = bytes()
EncMethod = ""


def LoadURI(u):

    if(re.match("http.*", u)):
        print(f'Loading URL : {u}')
        with urllib.request.urlopen(u) as f:
            F = f.read()
            print(f"Loaded {len(F)} bytes")
            return F
    else:
        print(f'Loading file : {u}')
        f = open(u, 'rb')
        F = f.read()
        return F


def ParseKey():
    global EncKey, EncKeyProvider, EncMethod

    if(EncMethod != ""):
        F = LoadURI(EncKeyProvider)
        EncKey = F
        print(f"Key parsed : Length {len(F)}")


def AskForKey():
    global EncKey, EncKeyProvider, EncMethod

    if(EncMethod != ""):
        r = input(
            "Key info found. Do you want to parse it? ( [y]es / [n]o - I'll provide a key / [i]gnore - Don't use keys ) : ")

        if(r.lower().startswith("y")):
            ParseKey()
        elif (r.lower().startswith("n")):
            EncMethodInput = input(
                f"Enter crypto algorithm (AES-128, AES-256,... ) . Enter nothing to use current ({EncMethod}) : ")
            if(len(EncMethodInput) != 0):
                EncMethod = EncMethodInput

            EncKeyProvider = input("Enter key provider. URL or Filename : ")
            if(len(EncKeyProvider) == 0):
                EncKeyProvider = wsdir + 'video.key'

            ParseKey()
        elif (r.lower().startswith("i")):
            EncMethod = ""
            EncKeyProvider = ""
            EncKey = []
        else:
            print("What?")
            AskForKey()
    else:
        r = input(
            "Key info not found. Do you want to use a key? ( [y]es / [N]o ) : ")

        if(r.lower().startswith("y")):
            EncMethod = input(
                "Enter crypto algorithm (AES-128, AES-256,... ) : ")
            EncKeyProvider = input("Enter key provider. URL or Filename : ")
            ParseKey()
        else:
            EncMethod = ""
            EncKeyProvider = ""


def main():
    global EncKey, EncKeyProvider, EncMethod

    print("M3U Downloader")

    filename = "input.json"

    if(len(argv) > 1):
        filename = argv[1]  # Read from args
    else:
        filename = input("Enter the M3U file name : ")
        if(len(filename) == 0):
            filename = f'{wsdir}playlist.m3u8'

    fn = open(filename, "rt")
    F = fn.read()

    EncKeys = re.findall("#EXT-X-KEY:METHOD=(.*),URI=(.*)", F)

    print(EncKeys)

    if(len(EncKeys) > 0):
        if(len(EncKeys[0]) == 2):
            EncMethod = EncKeys[0][0]
            EncKeyProvider = EncKeys[0][1]

    # AskForKey()

    for line in F.splitlines():
        if(line.startswith('#') != True):
            # print(line)
            filename = line
            if(re.match('.*/*.', line)):
                rfilename = re.findall("(?:.*\/)(.+)", line)
                if(len(rfilename) != 0):
                    filename = rfilename[0]

            Chunk = LoadURI(line)
            print(f"{filename} {len(Chunk)} bytes")
            FC = open(wsdir + filename, 'wb')
            FC.write(Chunk)

    # print(URIs)


main()

# LoadURI('https://example.com/')
