Magic Squares calculator on Go

This is my training project on Go, the algorithm uses the possibilities of productive and multi-flow recursive calculation of rearrangements of the source figures of the source array.

This is a light version (without BigInt), and it does not support squares with a dimension of more than 7

Symmetric checks for:<br/>
-horizontals<br/>
-verticals<br/>
-diagonals

Using:<br/>
-Set in the source file the dimension of the calculated square<br/>
-Run it

2x CPU Xeon 2689 v2, 2.6 MHz (20core/40threads)
3x3 squares perfomance<br/>
<img width="317" height="170" alt="image" src="https://github.com/user-attachments/assets/3199f8dd-3298-403e-bfcc-5d7a6b5e24c4" />


4x4 squares perfomance<br/>
<img width="425" height="132" alt="image" src="https://github.com/user-attachments/assets/a4fe7422-c881-48ac-a482-9e5c8282bc06" />

squares_old contains a version of the original number permutation algorithm that has poor performance

"draw" directory includes 4x4 squares in SVG images
<img width="480" height="480" alt="image" src="https://github.com/user-attachments/assets/4ec0e6b6-e9dd-4d9b-b30f-38af0db7b863" />
