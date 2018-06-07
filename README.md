ข้อตกลง commit message
[Create] : สร้างไฟล์
[Edit] : แก้ไข code ตัวเดิมที่มีอยู่แล้ว
[Add] : กรณีเพิ่มฟังก์ชันหรือ method
[Delete] : ลบไฟล์ ฟังก์ชัน หรือ method

ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน###
ตัวอย่างเช่น [Edit] : add parameter days in function setHours###

การกำหนดขั้นตอนก่อน code ในเครื่องหลัง pull ทุกครั้ง set ค่าใน command ดังนี้

export GOPATH=$(pwd)
go install packagename
go build -o ./bin/namefileinbin packagename

แต่เวลาที่จะ run ต้องเรียกใช้จาก main ดังนั้นต้องset "go build -o ./bin/main main" ทุกครั้ง

./bin/main เพื่อทดสอบดูผลลัพธ์