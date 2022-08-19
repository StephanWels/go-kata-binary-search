## Binary Search 
Create a function ```IndexOf``` that takes a sorted Array of Intergers and an Integer x as its arguments and returns the position of x
inside the Array.

## Description
Given a sorted array arr[] of n elements, write a function to search a given element x in arr[] and return the index of x in the array.
If the element can not be found, return -1.
The idea of binary search is to use the information that the array is sorted and reduce the time complexity to O(Log n).

## Example
Input:
arr[] = {10, 20, 30, 50, 60, 80, 110, 130, 140, 170}
x = 110

Output:
IndexOf(arr[], 110) -> 6

Output (not found):
IndexAOf(arr[], 5) -> -1