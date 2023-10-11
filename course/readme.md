
### Functions
1. Implement the function maxPower() that takes two parameters M and N and returns the largest integer K such that M <= NK
      maxPower (80000, 5) should return 8
      maxPower (30000, 9) should return 6
    Check correctness of the function in main().

### String operations
2. Implement the function num2words() which takes a number and returns the number in words as a string.
  For example, num2words(123) should return the string "one two three".
  Test correctness of the function with different argument, calling it from main().

3. Implement the function words2num() which takes a string and returns the number.
  For example, words2num("one two three") should return 123.
  Test correctness of the function with different argument, calling it from main().

4. Implement the function capitalizeWords() which takes a string of words and returns a string containing the words such that first letter of each word is capitalized.
For example, capitalizeWords ("one two three") should return "One Two Three".

5. Implement the function removeConsecutiveDuplicates() which takes two parameters, a string from which duplicates have to be removed and another string that specifies characters to be removed. The function should return the resulting string
removeConsecutiveDuplicates("gooddeed", "od") should return "geed"

### Arrays & Slices
6. Write a program that creates slices each of increasing size, starting with slice of 1 element.
    Store the first 120 prime numbers into these slices as per the pattern given below:
      Slice-1 : first prime number
      Slice-2 : next 2 prime numbers
      Slice-3 : next 3 prime numbers
      Slice-4 : next 4 prime numbers
    Having created, print contents each slice one line per row.

### Commandline arguments
7. Implement the function num2hex() which takes an unsigned integer and returns its hexa equivalent as a string.
Function should take an extra argument which specifies whether the hexa digits should be in lower case (default) or upper case.
Using the function num2hex() implement main which takes an integer value and prints its hexa equivalent. If the number precedes with —u option, then the hexa value should be printed in uppercase. Any other option other than —u should be considered invalid.

8. Write a program that takes 2 complex numbers and operation to be performed on them as command line arguments, and prints the result as complex number.
  Commandline arguments to the program are in the following format
    complexnum1 binaryop complexnum2
  Note: Assume there is no embedded space in complex number argument.
    ex:
    3-4i + 7+2i
    -15i - 3+4i
    3-5i * 12


### Pointers & structures
9. Define structure EmpInfo that holds the following
    Empno(5 digits),
    date-of-join (dd, mm, yy),
    locationcode (string of 4 characters)
  Create an array of 10 elements of the structure type and initialize the structure, in random order.
  Implement the functions
    - printEmpByLocation() which takes pointer to EmpInfo structure & location code as arguments, and prints list of employees & their date of join.
    - printEmpByYearOfJoin() which takes pointer to EmpInfo structure & year as arguments, and prints list of employees, date of join & location.
  The program should support the following command line agruments:
    -I locationcode → should print employees of the specified location
    -y year →should print employees in that particular year

### Interfaces
10. Create a structure named country_cap with members country & capital city.
      type country_cap struct {
        country string
        population int
      }
    Declare an array of 10 elements and initialize it with data.
      a) Sort the array by country name in ascending order and print the elements.
      b) Sort the array by population in descending order and print the elements.
    Hint: use Sort function of the package "sort" & implement sort. Interface for each sorting criteria.

### File I/O
11. Implement the function splitCopy() as per the details given below:
      a) Function takes source file handle, destination file handle1, destination file handle2 
        i) copies odd numbered lines to the file whose file handle is destination file handle1 
        ii) copies even numbered lines to the file whose file handle is destination file handle2
      b) Use the above function to copy odd and even numbered lines of file X into file Y and Z respectively. Filenames X, Y and Z are given as command line arguments.

### Maps & File I/O
12. Write a program that reads words from console, and at the end of the input, prints words of same length. Output only those words which have at least 4 characters.
    Output should be in the following format:
      4: 4-letter words in alphabetical order
      5: 5-letter words in alphabetical order
      6: 6-letter words in alphabetical order
      ...