# Lexicon

Examples:
- `lexicon create resource -f file.yaml`
- ```
   cat <<EOF | lexicon create resource -f -
    name: foo
    labels:
     - alpha
     - bravo
    EOF
    ```
- `lexicon delete resource --identifier={name=foo}`
