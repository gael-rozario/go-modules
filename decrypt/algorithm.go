//Package for decryption algorithms using simple symmetric ceaser cypher
package decrypt

//Function for decryption algorithm
func Decrypt(input string) string{
    decryptedString := ""
    for _,ch := range input{
        asciiCode := int(ch)
        decryptedChar := string(asciiCode - 3)
        decryptedString += decryptedChar
    }
    return decryptedString 
}
