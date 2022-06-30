use aragog::Record;
use derive_builder::Builder;
use serde::{Deserialize, Serialize};

#[derive(Debug)]
enum Results {
    Lost(String),
    Won(String),
    Drop(String),
}

#[derive(Clone, Debug, Serialize, Deserialize, Builder)]
#[serde(rename_all = "lowercase")]
pub struct Player {
    pub _key: String,
    pub playerid: String,
    pub birthdate: String,
}

#[derive(Clone, Debug, Serialize, Deserialize, Record, Builder)]
#[serde(rename_all = "lowercase")]
pub struct Venue {
    pub _key: String,
    pub address: String,
    pub lat: String,
    pub lng: String,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct VenueRequest {
    pub name: String,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct VenueResponse {
    pub address: String,
    pub name: String,
}

impl VenueResponse {
    pub fn of(venue: Venue) -> VenueResponse {
        VenueResponse {
            address: venue.address,
            name: venue.lat,
        }
    }
}

#[derive(Clone, Debug, Serialize, Deserialize)]
#[serde(rename_all = "lowercase")]
pub struct Game {
    pub _key: String,
    pub name: String,
}

#[derive(Clone, Debug, Serialize, Deserialize, Builder)]
#[serde(rename_all = "lowercase")]
#[builder(setter(into))]
pub struct ContestGraph {
    pub _key: String,
    pub start: String,
    pub startoffset: String,
    pub stop: String,
    pub stopoffset: String,
}

#[derive(Clone, Debug, Serialize, Deserialize, Builder)]
#[serde(rename_all = "lowercase")]
#[builder(setter(into))]
pub struct Contest {
    pub contest_graph: ContestGraph,
    pub venue: PlayedAt,
    pub outcome: Vec<ResultedIn>,
    pub games: Vec<PlayedWith>,
}

#[derive(Clone, Debug, Serialize, Deserialize, Builder)]
#[serde(rename_all = "lowercase")]
#[builder(setter(into))]
pub struct PlayedAt {
    pub _key: String,
    pub _to: String,
    pub _from: String,
    pub _label: String,
}

#[derive(Clone, Debug, Serialize, Deserialize, Builder)]
#[serde(rename_all = "lowercase")]
#[builder(setter(into))]
pub struct ResultedIn {
    pub _key: String,
    pub _to: String,
    pub _from: String,
    pub _label: String,
    pub place: i32,
    pub results: String,
}

#[derive(Clone, Debug, Serialize, Deserialize, Builder)]
#[serde(rename_all = "lowercase")]
#[builder(setter(into))]
pub struct PlayedWith {
    pub _key: String,
    pub _to: String,
    pub _from: String,
    pub _label: String,
}

//**** IMPORT FROM JSON OBJECTS */
#[derive(Clone, Debug, Serialize, Deserialize, Builder)]
#[serde(rename_all = "lowercase")]
pub struct ContestImport {
    pub _key: String,
    pub start: String,
    pub startoffset: String,
    pub stop: String,
    pub stopoffset: String,
    pub venue: Venue,
    pub outcome: Vec<Outcome>,
    pub games: Vec<String>,
}
#[derive(Clone, Debug, Serialize, Deserialize)]
#[serde(rename_all = "lowercase")]
pub struct Outcome {
    pub _key: String,
    pub playerid: String,
    pub place: String,
    pub result: String,
}

/* OLDE STUFF TO BE YANKED */
#[derive(Deserialize, Clone, PartialEq, Debug)]
pub struct Owner {
    pub id: i32,
    pub name: String,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct OwnerRequest {
    pub name: String,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct OwnerResponse {
    pub id: i32,
    pub name: String,
}

impl OwnerResponse {
    pub fn of(owner: Owner) -> OwnerResponse {
        OwnerResponse {
            id: owner.id,
            name: owner.name,
        }
    }
}

#[derive(Deserialize, Clone, PartialEq, Debug)]
pub struct Pet {
    pub id: i32,
    pub name: String,
    pub owner_id: i32,
    pub animal_type: String,
    pub color: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct PetRequest {
    pub name: String,
    pub animal_type: String,
    pub color: Option<String>,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
pub struct PetResponse {
    pub id: i32,
    pub name: String,
    pub animal_type: String,
    pub color: Option<String>,
}

impl PetResponse {
    pub fn of(pet: Pet) -> PetResponse {
        PetResponse {
            id: pet.id,
            name: pet.name,
            animal_type: pet.animal_type,
            color: pet.color,
        }
    }
}
