# <span style="color:#C0BFEC">**🦔 Utility to find unique strings**</span>

## <span style="color:#C0BFEC">***Enter to run:*** </span>

```
go run uniq.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]
```

## <span style="color:#C0BFEC">***Options:***</span>

`-c` - count the number of occurrences of the string in the input.
Output this number before the string separated by a space.

`-d` - output only those lines that are repeated in the input.

`-u` - output only those lines that are not repeated in the input.

`-f num_fields` Ignore the first `num_fields` fields in a line.
A field in a string is a non-empty set of characters separated by a space.

`-s num_chars` ignore the first `num_chars` characters in the string.
When used with the `-f` option, first characters are counted
after `num_fields` fields (ignoring space delimiter after
last field).

`-i` - do not take into account the case of letters.

## <span style="color:#C0BFEC">***Usage:***</span>

`uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]`

1. All parameters are optional. Utility behaviors without parameters --
   simple derivation of unique strings from the input.

2. Parameters `c, d, u` are interchangeable. Should be considered,
   that in parallel these parameters do not make any sense. At
   passing one along with the other needs to be displayed to the user
   proper use of the utility

3. If `input_file` is not passed, then consider `stdin` as the input stream

4. If `output_file` is not passed, then consider `stdout` as the output stream

## <span style="color:#C0BFEC">***Examples:***</span>

<details>
    <summary>Without parameters</summary>

```shell
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go
I love music.

I love music of Kartik.
Thanks.
```

</details>

<details>
    <summary>With "input_file" parameter</summary>

```shell
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
$go run uniq.go input.txt
I love music.

I love music of Kartik.
Thanks.
```

</details>

<details>
    <summary>With "input_file" and "output_file" parameters</summary>

```shell
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
$go run uniq.go input.txt output.txt
$cat output.txt
I love music.

I love music of Kartik.
Thanks.
```

</details>

<details>
    <summary>With "-c" parameter</summary>

```shell
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go -c
3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
```

</details>

<details>
    <summary>With "-d" parameter</summary>

```shell
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go -d
I love music.
I love music of Kartik.
```

</details>

<details>
    <summary>With "-u" parameter</summary>

```shell
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go -d

Thanks.
```

</details>

<details>
    <summary>With "-i" parameter</summary>

```shell
$cat input.txt
I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
$cat input.txt | go run uniq.go -i
I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
```

</details>

<details>
    <summary>With "-f num" parameter</summary>

```shell
$cat input.txt
We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go -f 1
We love music.

I love music of Kartik.
Thanks.
```

</details>

<details>
    <summary>With "-s num" parameter</summary>

```shell
$cat input.txt
I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
$cat input.txt | go run uniq.go -s 1
I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
```

</details>

