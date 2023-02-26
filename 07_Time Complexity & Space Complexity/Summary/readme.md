# Time Complexity

## Pengertian
Time Complexity adalah algoritma yang digunakan untuk memperkirakan waktu yang akan dihabiskan oleh sebuah program untuk menyelesaikan sebuah masalah dengan input tertentu. Biasanya dinotasikan dengan menggunakan *Big O Notation*. Notasi ini memberikan batas atas waktu yang diperlukan oleh sebuah komputer apabila input menjadi sangat besar.

## Jenis Big O Notation

### O(1)
O(1) adalah sebuah Notasi O yang konstan maupun linear. Biasanya hal ini terjadi karena tidak ada batasan perulangan sebanyak n buah maupun sebuah rekursif. Contoh dari O(1) adalah :
```
result += result + 1
```

### O(n)
O(n) adalah sebuah Notasi O yang biasanya berada pada sebuah perulangan, dimana perulangan tersebut dibatasi hingga sebuah nilai. Contoh dari O(n) adalah: 
```
for i:=0; i<=n; i++{
    result += result + i
}
```

### O(n^a)
O(n^a) adalah sebuah notasi O yang biasanya menunjukkan seberapa banyaknya *Nested Loop* dalam sebuah program. Biasanya nilai a adalah banyaknya *Nested Loop*. Contoh dari O(n^2) adalah:
```
for i<=n{
    for j<=n{
        result += i*j
    }
}
```

### O(log n)
O(log n) adalah sebuah Notasi O yang biasanya terdapat dalam sebuah perulangan seperti O(n), akan tetapi seiring berjalannya perulangan tadi, nilai dari batas perulangan tadi akan menjadi setengah dari sebelumnya. Contoh dari O(log n):
```
for i:=0; i<=n; i++{
    result += result + i
    n=n/2
}