import { Registred } from "./registred.js";
import { Get_Data } from "./fetch_data.js";
import { F } from "./fetch_data.js";
export async function filter() {
  let filterbutton = document.querySelector(".filterbutton");

  filterbutton.onclick = async function (event) {
    event.preventDefault();

    let postsData = await Get_Data(`/api`);

    postsData = postsData.posts;
    let filtered = postsData;

    let selectFilter = document.querySelector(".selectfilter");

    for (const e of selectFilter.selectedOptions) {

      if (e.value == "sport") {
        filtered = sportFilter(filtered);
      } else if (e.value == "science") {
        filtered = scienceFilter(filtered);
      } else if (e.value == "entertainment") {
        filtered = enetertimentFilter(filtered);
      } else if (e.value == "created") {
        filtered = await createdFilter(filtered);
      } else if (e.value == "liked") {
        filtered = await likedFilter(filtered);
      }
    }



    F(filtered)
  };
}

function sportFilter(postsData) {
  let filtered = [];
  for (let i = 0; i < postsData.length; i++) {
    if (Array.isArray(postsData[i].categories)) {
      if (postsData[i].categories.includes("sport")) {
        filtered.push(postsData[i]);
      }
    }
  }

  return filtered;
}

function scienceFilter(postsData) {
  let filtered = [];
  for (let i = 0; i < postsData.length; i++) {
    if (Array.isArray(postsData[i].categories)) {
      if (postsData[i].categories.includes("science")) {
        filtered.push(postsData[i]);
      }
    }
  }

  return filtered;
}

function enetertimentFilter(postsData) {
  let filtered = [];
  for (let i = 0; i < postsData.length; i++) {
    if (Array.isArray(postsData[i].categories)) {
      if (postsData[i].categories.includes("entertainment")) {
        filtered.push(postsData[i]);
      }
    }
  }


  return filtered;
}

async function likedFilter(postsData) {
  let filtered = [];

  let userid = await Registred();

  if (!userid) {
    window.location.replace("/login");
  } else {


    for (let i = 0; i < postsData.length; i++) {
      if (Array.isArray(postsData[i].likers)) {

        if (postsData[i].likers.includes(userid)) {

          filtered.push(postsData[i]);
        }
      }
    }

    return filtered;
  }
}

async function createdFilter(postsData) {
  let filtered = [];

  let id = await Registred()
  if (!id) {
    window.location.replace("/login");
  } else {


    for (let i = 0; i < postsData.length; i++) {

      if (postsData[i].user_id === id) {
        filtered.push(postsData[i]);
      }

    }
  }


  return filtered;
}
