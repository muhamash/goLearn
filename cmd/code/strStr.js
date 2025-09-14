
function strStr(haystack, needle) {
    if ( haystack.length === 0 || needle.length === 0 )
    {
        return false
    }

    // return haystack.indexOf( needle )

    for ( let i = 0; i <= haystack.length - needle.length - 1; i++ )
    {
        let j = 0;
        while ( j < haystack.length && haystack[ i + j ] === needle[ j ] )
        {
            j++
        }

        if ( j === needle.length )
        {
            return i
        }


        return -1
    }
    

}

console.log(strStr("sadbutsad", "sad"));
console.log(strStr("leetcode", "leeto")); 