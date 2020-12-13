'use strict';
const fetch = require('node-fetch');

const Client = (serverURL) => {
    const respHandler = (resp) => {
        if (resp.ok) {
            return resp.json();
        }
        throw new Error(`Unexpected response from the server ${resp.status} ${resp.statusText}`);
    };

    return {
        get: async (path) => {
            return await fetch(`${serverURL}${path}`);
        },
        post: async (path, data) => {
            return await fetch(`${serverURL}${path}`, {
                method: 'post',
                body: JSON.stringify(data),
                headers: {'Content-Type': 'application/json'}
            });
        }
    };
};

module.exports = { Client };
