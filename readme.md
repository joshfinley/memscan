# memscan

Toy go memory mapping tool

## Example

```
 go run .\cmd\main.go -p 7380
+-----------------------------------------------------------------------------------------------------------+
| Memory Map for Process 7380                                                                               |
+----------------+------------+---------+-------------------+-----------------------------------------------+
| ADDRESS        | SIZE       | TYPE    | PROTECT           | MODULE PATH                                   |
+----------------+------------+---------+-------------------+-----------------------------------------------+
| 0x7ff91835c000 | 0x2000     | Image   | PAGE_READWRITE    |                                               |
| 0x7ff920225000 | 0x11000    | Image   | PAGE_READONLY     |                                               |
| 0x7ff91ed3a000 | 0x4000     | Image   | PAGE_READONLY     |                                               |
| 0x7ff5770a4000 | 0x3000     | Mapped  | PAGE_READONLY     |                                               |
| 0x7ff91e6e6000 | 0x1000     | Image   | PAGE_READWRITE    |                                               |
| 0x7df5928b1000 | 0xf000     |         | PAGE_NOACCESS     |                                               |
| 0x7ff57708c000 | 0x18000    | Mapped  | PAGE_NOACCESS     |                                               |
| 0x7df5946ad000 | 0xfed25000 | Mapped  |                   |                                               |
| 0x7ff920050000 | 0x1000     | Image   | PAGE_READONLY     | C:\WINDOWS\System32\bcryptPrimitives.dll      |
| 0x7ff917525000 | 0x2000     | Image   | PAGE_READWRITE    |                                               |
| 0x7ff922362000 | 0x86000    | Image   | PAGE_READONLY     |                                               |
| 0x7ff91fed3000 | 0x4000     | Image   | PAGE_READWRITE    |                                               |
| 0x7ff91f7b0000 | 0x1000     | Image   | PAGE_READONLY     | C:\WINDOWS\SYSTEM32\SSPICLI.DLL               |
| ................................................ omitted ................................................ |
+----------------+------------+---------+-------------------+-----------------------------------------------+
```