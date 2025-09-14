
function strStr(haystack, needle) {
    if ( haystack.length === 0 || needle.length === 0 )
    {
        return false
    }

    return haystack.indexOf(needle)
}

console.log(strStr("sadbutsad", "sad"));
console.log(strStr("leetcode", "leeto")); 