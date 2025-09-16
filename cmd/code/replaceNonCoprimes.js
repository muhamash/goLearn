
// helper ot find gcd
function gcd ( a, b )
{
    while ( b !== 0 )
    {
        let temp = b
        b = a % b
        a = temp
    }

    return a
}

// helper to find lcm
function lcm ( a, b )
{
    return ( a * b ) / gcd( a, b );
}

// main function
function replaceNonCoprimes ( nums )
{
    const stack = []
    
    for ( let num of nums )
    {
        while ( stack.length > 0 )
        {
            const top = stack[ stack.length - 1 ];
            const g = gcd( top, num )
            if ( g > 1 )
            {
                num = lcm( top, num )
                stack.pop()
            }
            else
            {
                break;
            }
        }

        stack.push(num)
    }

    return stack
}

console.log(replaceNonCoprimes([6, 4, 3, 2, 7, 6, 2])); // [12, 7, 6]
console.log(replaceNonCoprimes([2, 2, 1, 1, 3, 3, 3])); // [2, 1, 1, 3]