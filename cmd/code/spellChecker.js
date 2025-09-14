
function spellChecker (wordlist, queries)
{
    // helper to check vowels
    const vowels = new Set( [ 'a', 'e', 'i', 'o', 'u' ] )
    const devowel = ( word ) => {
        word.toLowerCase().replace( /[aeiou]/g, '*' );
    }


    // exact word set
    const exactWords = new Set( wordlist )
    
    // case insensitive map
    const caseInsensitive = new Map()

    // vowel error map
    const vowelInsensitive = new Map()

    for ( const word of wordlist )
    {
        const lower = word.toLowerCase()
        if ( !caseInsensitive.has( lower ) )
        {
            caseInsensitive.set(lower, word)
        }

        const dev = devowel( word );
        if ( !vowelInsensitive.has( dev ) )
        {
            vowelInsensitive.set(dev, word)
        }

    }

    const result = []
    for ( const q of queries )
    {
        // exact match
        if ( exactWords.has( q ) )
        {
            result.push( q )
            continue
        }

        const lower = q.toLowerCase()

        // case insensitive
        if ( caseInsensitive.has( lower ) )
        {
            result.push( caseInsensitive.get( lower ) )
            continue
        }

        // vowel insensitive
        const dev = devowel( q )
        if ( vowelInsensitive.has( dev ) )
        {
            result.push( vowelInsensitive.get( dev ) )
            continue
        }

        // no match
        result.push("")
    }

    return result
}

const wordlist1 = ["KiTe","kite","hare","Hare"];
const queries1 = ["kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"];
console.log(spellChecker(wordlist1, queries1));


const wordlist2 = ["yellowfds","aeiouasd", "AEUIsadOI", "fdsf", "aesdfjkg"];
const queries2 = ["YellOw", "aeiou", "AEUIOI"];
console.log(spellChecker(wordlist2, queries2));