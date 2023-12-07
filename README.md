# Enable GCC
```
go env CGO_ENABLED //if 0 set to 1
export CGO_ENABLED=1
```

# Ubuntu
```
sudo apt install build-essential libtesseract-dev
wget https://github.com/tesseract-ocr/tessdata/raw/main/eng.traineddata
sudo mv eng.traineddata /usr/share/tesseract-ocr/4.00/tessdata/eng.traineddata
```
