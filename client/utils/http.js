'use strict';
const fetch = require('node-fetch');

const Client = (serverURL) => {
    const respHandler = (resp) => {
        if (resp.status === 200) {
            return resp.json();
        } else if (resp.status === 400) {
            console.log(`${resp.status}: ${resp.statusText}`);
            return resp.json();
        } else {
            throw new Error(`Unexpected response from server ${resp.status}: ${resp.statusText}`);
        }
    };

    return {
        get:  (path) => {
            return fetch(`${serverURL}${path}`)
                .then(res => respHandler(res))
                .catch(err => console.log(err));
        },
        post: (path, data) => {
            return fetch(`${serverURL}${path}`, {
                method: 'post',
                body: JSON.stringify(data),
                headers: {'Content-Type': 'application/json'}
            })
                .then(res => respHandler(res))
                .catch(err => console.log(err));
        }
    };
};

module.exports = { Client };
