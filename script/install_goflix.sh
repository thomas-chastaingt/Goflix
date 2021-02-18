#bin/sh

echo " ---------- Start deploying Goflix ---------- "
cd ../
go build -o goflix_64_macos
mv ./goflix_64_macos ./bin
echo " ---------- Goflix deployed enjoy ! ---------- " 
