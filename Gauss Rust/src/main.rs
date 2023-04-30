use std::{time::{Instant, Duration}, ops::{Index, IndexMut}, borrow::Borrow};
use rand::{Rng, rngs::StdRng, SeedableRng};

fn main() {
    // Ordem da matriz
    let n = 2000;

    // let (mut A, mut B, mut X) = create_values(n);
    let (mut A, mut B, mut X) = create_values_array2d(n);

    let before = Instant::now();
    // gauss_solver(&mut A, &mut B, &mut X);
    gauss_solver_Array2D(&mut A, &mut B, &mut X);
    let now = Instant::now();

    println!("X: {:?}", X);
    show_time((now - before).borrow());
}

fn show_time(duration: &Duration) {
    let ms = duration.as_millis() % 1000;
    let s = duration.as_secs() % 60;
    let m = duration.as_secs() / 60 % 60;
    let h = duration.as_secs() / 60 / 60;
    println!("Total time: {}s", duration.as_secs_f64());
    println!("\t{h:02}h:{m:02}m:{s:02}s:{ms:03}ms"); 
}

fn create_values(n: usize) -> (Vec<Vec<f64>>, Vec<f64>, Vec<f64>) {
    // let mut random = thread_rng();
    let mut r = StdRng::seed_from_u64(0);   // Reproducible random sequence
    let mut A: Vec<Vec<f64>> = Vec::new();
    let B: Vec<f64> = vec![1.0; n];
    let X: Vec<f64> = vec![0.0; n];

    for _ in 0..n {
        let mut temp = vec![0.0; n];
        r.try_fill(&mut temp[..]).unwrap();
        A.push(temp);
    }

    return (A, B, X);
}

fn gauss_solver(A: &mut Vec<Vec<f64>>, B: &mut Vec<f64>, X: &mut Vec<f64>) {
    let N = A.len();

    for norm in 0..(N - 1) {
        for row in (norm + 1)..N {
            let multiplier = A[row][norm] / A[norm][norm];
            for col in norm..N {
                A[row][col] -= A[norm][col] * multiplier;
            }
            B[row] -= B[norm] * multiplier;
        }
    }

    for row in (0..N).rev() {
        X[row] = B[row];
        for col in ((row + 1)..N).rev() {
            X[row] -= A[row][col] * X[col];
        }
        X[row] /= A[row][row];
    }
}


fn create_values_array2d(n: usize) -> (Array2D<f64>, Vec<f64>, Vec<f64>) {
    let mut r = StdRng::seed_from_u64(0);   // Reproducible random sequence
    let mut A = Array2D::new(n, n);
    let B: Vec<f64> = vec![1.0; n];
    let X: Vec<f64> = vec![0.0; n];

    for index in 0..n {
        r.try_fill(&mut A[index]).unwrap();
    }

    return (A, B, X);
}

fn gauss_solver_Array2D(A: &mut Array2D<f64>, B: &mut Vec<f64>, X: &mut Vec<f64>) {
    let N = A.columns;

    for norm in 0..(N - 1) {
        for row in (norm + 1)..N {
            let multiplier = A[row][norm] / A[norm][norm];
            for col in norm..N {
                A[row][col] -= A[norm][col] * multiplier;
            }
            B[row] -= B[norm] * multiplier;
        }
    }

    for row in (0..N).rev() {
        X[row] = B[row];
        for col in ((row + 1)..N).rev() {
            X[row] -= A[row][col] * X[col];
        }
        X[row] /= A[row][row];
    }
}

pub struct Array2D<T> {
    vec: Vec<T>,
    columns: usize,
    rows: usize,
}

impl<T> Array2D<T> {
    pub fn new(rows: usize, columns: usize) -> Self
    where
        T: Default,
        T: Clone,
    {
        let mut new = Array2D {
            vec: Vec::with_capacity(columns * rows),
            columns: columns,
            rows: rows,
        };
        for _ in 0..(columns * rows) {
            new.vec.push(T::default());
        }
        return new;
    }

    pub fn rows_len(&self) -> usize {
        return self.rows;
    }

    pub fn columns_len(&self) -> usize {
        return self.columns;
    }
}

impl<T> Index<usize> for Array2D<T> {
    type Output = [T];

    fn index(&self, index: usize) -> &Self::Output {
        let index_start = index * self.columns;
        let index_end = (index + 1) * self.columns;

        &self.vec[index_start..index_end]
    }
}

impl<T> IndexMut<usize> for Array2D<T> {
    fn index_mut(&mut self, index: usize) -> &mut Self::Output {
        let index_start = index * self.columns;
        let index_end = (index + 1) * self.columns;

        &mut self.vec[index_start..index_end]
    }
}
