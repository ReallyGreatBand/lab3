'use strict';
const http = require('../utils/http');

const Client = (serverURL) => {
    const client = http.Client(serverURL);

    return {
        listPlants: () => client.get('/plants'),
        updatePlant: (id, soilMoistureLevel) => client.post('/plants', { id, soilMoistureLevel })
    }
};

module.exports = { Client };
