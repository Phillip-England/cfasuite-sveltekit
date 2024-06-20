
import { SERVER_URL } from "$lib/config/config";

export class Request {


    static async post(path, data) {
        return await fetch(`${SERVER_URL}${path}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
            credentials: 'include',
        });
    }

    static async get(path, data) {
        return await fetch(`${SERVER_URL}${path}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
        });
    }

}