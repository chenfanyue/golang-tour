// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

for key, value := range oldMap {
    newMap[key] = value
}

for key := range m {
    if key.expired() {
        delete(m, key)
    }
}

sum := 0
for _, value := range array {
    sum += value
}

for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
// character U+65E5 '日' starts at byte position 0
// character U+672C '本' starts at byte position 3
// character U+FFFD '�' starts at byte position 6
// character U+8A9E '語' starts at byte position 7

// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}