'use strict';

const plants = require('./plants/client');

const client = plants.Client('http://localhost:8080');

const plantsList = (await client.listPlants()).json;
console.log('Show plants with soil moisture level less than 0.2');
console.log(plantsList);

const updatedPlant = (await client.updatePlant(1, 1)).json;
console.log('Update moisture level for plant with id');
console.log(updatedPlant);

