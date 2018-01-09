# FastqCount

**Usage**:

summary single fastq(.gz) file:

    ```./FastqCount input.fastq```

multipy files:
  ```pigz -dc R1.fastq.gz R2.fastq.gz | ./FastqCount -```

Note: pipeline stdin runs faster for FastqCount with gzipped file(s).
