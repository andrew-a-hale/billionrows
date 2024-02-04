use io::BufRead;
use std::{collections::HashMap, env, fs::File, io, path::Path};

#[derive(Clone, Debug)]
struct State {
    map: HashMap<String, Vec<f32>>,
    size: i64,
}

impl State {
    fn new() -> Self {
        State {
            map: HashMap::new(),
            size: 0,
        }
    }

    fn process(&mut self, row: String) -> Self {
        let mut values = row.split(';');
        let station = match values.next() {
            Some(x) => String::from(x),
            None => panic!("Malformed data: {}", row),
        };
        let value = match values.next() {
            Some(x) => x.parse::<f32>().unwrap(),
            None => panic!("Malformed data: {}", row),
        };

        self.size += 1;
        let mut vec = vec![value, value, value];
        match self.map.get(&station) {
            Some(x) if &value < x.first().unwrap() => {
                vec[1] += x[1];
                vec[2] = x[2];
            }
            Some(x) if &value > x.last().unwrap() => {
                vec[0] = x[0];
                vec[1] += x[1];
            }
            Some(x) => {
                vec[0] = x[0];
                vec[1] += x[1];
                vec[2] = x[2];
            }
            None => (),
        }

        self.map.insert(station, vec);
        self.clone()
    }
}

fn main() {
    // scan file with accumulator
    let args: Vec<String> = env::args().collect();
    let filepath = args.get(1).expect("Expected filepath");

    let state = read_lines(filepath)
        .unwrap()
        .scan(State::new(), |state, x| {
            state.process(x.unwrap());
            Some(state.clone())
        })
        .last()
        .unwrap();

    for (station, mut aggs) in state.map {
        aggs.insert(aggs.len(), aggs.get(1).unwrap() / state.size as f32);
        println!("{}: {:?}", station, aggs)
    }
}

fn read_lines<P>(filepath: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filepath)?;
    Ok(io::BufReader::new(file).lines())
}
