'use strict';
const http = require('../utils/http');

const Client = (serverURL) => {
    const client = http.Client(serverURL);

    return {
        listPlants: async () => await client.get('/plants'),
        updatePlant: async (id, soilMoistureLevel) => await client.post('/plants', { id, soilMoistureLevel })
    }
};

module.exports = { Client };
