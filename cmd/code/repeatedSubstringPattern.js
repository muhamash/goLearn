
function repeatedSubstringPattern ( s )
{
    const doubleString = s + s;

    return doubleString.slice(1, doubleString.length - 1).includes(s)
}

console.log(repeatedSubstringPattern("abab"));
console.log(repeatedSubstringPattern("aba")); 
console.log(repeatedSubstringPattern("abcabcabcabc"));