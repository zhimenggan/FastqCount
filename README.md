# FastqCount

Summary fastq reads, bases, N bases, Q20, Q30, GC.

**Usage**:

summary single fastq(.gz) file:

```./FastqCount input.fastq```

summary multipy files:

```pigz -dc R1.fastq.gz R2.fastq.gz | ./FastqCount -```

**Note**: pipeline stdin make FastqCount faster with gzipped file(s).

**Output** example (tsv):

| Total Reads | Total Bases | N Bases | Q20 | Q30 | GC |
| ----------- | ----------- | ------- | --- | --- | -- |
| 8781961 (8.78 M) | 1317294150 (1.32 G) | 0.00% | 72.00% | 62.00% | 45.00% |
