sudo apt-get install zlib1g-dev
sudo apt-get install libpng12-dev libjpeg62-dev libtiff4-dev libsdl-gfx1.2-dev libsdl-image1.2-dev libsdl1.2-dev libavcodec-dev libavdevice-dev libavformat-dev libavutil-dev
sudo apt-get install python-enchant python-poppler
n_processors="$(grep '^processor' /proc/cpuinfo | wc -l)"

export LEPTONICA_VERSION="1.74"
cd /tmp
wget http://www.leptonica.com/source/leptonica-${LEPTONICA_VERSION}.tar.gz
tar -xvzof leptonica-${LEPTONICA_VERSION}.tar.gz
cd leptonica-${LEPTONICA_VERSION}
./configure
make -j${n_processors}
sudo make install
sudo ldconfig

export TESSERACT_VERSION="3.02.02"
cd /tmp
wget https://github.com/tesseract-ocr/tesseract/archive/${TESSERACT_VERSION}.tar.gz
tar -xvzof ${TESSERACT_VERSION}.tar.gz
cd tesseract-${TESSERACT_VERSION}
./autogen.sh
./configure
sudo rm -f /usr/local/share/tessdata
sudo mkdir /usr/local/share/tessdata
make -j${n_processors}
sudo make install
sudo ldconfig

cd /tmp
wget https://github.com/tesseract-ocr/tessdata/raw/3.04.00/eng.traineddata
wget https://github.com/tesseract-ocr/tessdata/blob/master/jpn.traineddata
cd /usr/local/share/tessdata
sudo cp /tmp/eng.traineddata /usr/local/share/tessdata
sudo cp /tmp/jpn.traineddata /usr/local/share/tessdata
