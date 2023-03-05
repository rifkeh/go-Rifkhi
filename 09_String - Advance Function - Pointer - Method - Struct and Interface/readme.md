# String - Advance Function - Pointer - Method - Struct and Interface

## String
String adalah tipe data yang biasanya mengandung teks. Dalam Go, untuk menginialisasi sebuah string, diperlukan tanda petik dua (" "). Dalam mengolah atau memainkan string, kita dapat dibantu dengan salah satu libray string, yaitu strings.

## Advance Function
Lebih dalam dari fungsi sebelumnya, advance function adalah sebuah function yang parameternya dapat berupa fungsi dan outputnya adalah fungsi. 

Biasanya kita memakai advance function ketika :
    - Melakukan fungsi callback
    - Melakukan fungsi generator
    - Melakukan fungsi closure

## Pointer
Pointer adalah cara untuk menunjuk ke address suatu variabel. Hal ini diperlukan karena biasanya dalam menginiliasikan sebuah variabel, contoh:
```
a := 10
b = a
```
Apabila kita mengganti nilai b, maka nilai a tidak akan berubah, hal ini karena ketika kita menginiliasisasikan a = b, maka hanya pass by value, atau hanya nilai saja yang berpindah, tidak ada kaitan lebih lanjut. Tapi jika kita,
```
a := 10
*b = &a
```
Maka apabila kita mengganti nilai b, nilai a akan mengikuti, begitu juga sebaliknya.

## Method dan Struct
Struct adalah salah satu tipe data  yang dapat membuat kumpulan data yang berbeda. Sedangkan method, adalah sebuah fungsi yang dapat dikaitkan dengan salah satu tipe data saja. Hal ini membuat method dapat memberi perilaku pada fungsi tersebut.