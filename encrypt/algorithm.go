package encrypt

func Encrypt(input string) string{
    encryptedString := ""
    for _, char:= range input{
        asciiCode := int(char)
        ecryptedChar := string(asciiCode+3)
        encryptedString += ecryptedChar
    }
    return encryptedString
}
