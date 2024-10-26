package decrypt

func Decrypt(input string) string{
    decryptedString := ""
    for _,ch := range input{
        asciiCode := int(ch)
        decryptedChar := string(asciiCode - 3)
        decryptedString += decryptedChar
    }
    return decryptedString 
}
