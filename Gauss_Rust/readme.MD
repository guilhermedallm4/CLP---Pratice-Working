## Implementação em Rust do método da Eliminação de Gauss
Existe duas versões:
1. Serial com representação da matriz como um vetor de vetores `Vec<Vec<f32>>`<br>
2. Serial com representação da matriz mapeada em um vetor, utilizando a forma [row major](https://en.wikipedia.org/wiki/Row-_and_column-major_order).<br>

Não existe implementado nenhuma interface para selecionar de forma simplificada qual algoritmo será executado, então é necessário descomentar a opção desejada na função `main` e recompilar e rodar :(

# Build & Run
Para executar, é necessário ter o compilador `Rust`. <br>
1. Instale o Rust utilizando o `rustup` de acordo com o [site oficial](https://www.rust-lang.org/tools/install)
2. Clone o repositório
```
git clone https://github.com/guilhermedallm4/CLP-Pratice-Working
cd 'CLP-Pratice-Working/Gauss_Rust'
```
3. Rode :)
```
cargo run
```
4. Se quiser compilar e executar de forma otimizada
```
cargo run --release
```