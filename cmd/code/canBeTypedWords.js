var canBeTypedWords = function ( text, brokenLetters )
{
    const brokenSet = new Set( brokenLetters ); 
    let count = 0;

    for ( const word of text.split( " " ) )
    {
        let canType = true;
        for ( const char of word )
        {
            if ( brokenSet.has( char ) )
            {
                canType = false;
                break;
            }
        }
        if ( canType ) count++;
    }

    return count;
};