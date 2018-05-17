def to_rna(dna_strand):
    dna_to_rna = {'G':'C','C':'G','T':'A','A':'U'} 
    rna = ""
    for c in dna_strand:
        try:
            rna += dna_to_rna[c] 
        except:
            raise ValueError("Given DNS strand is not valid!") 
    return rna
        
