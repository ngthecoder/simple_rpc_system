import * as net from "net"
import * as readline from "readline"

interface RPCRequest {
    method: string,
    params: any[],
    param_types: string[],
    id: number
}

interface RPCResponse {
    result: string,
    result_type: string,
    id: number
}

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
})

rl.question('Enter method name: ', (method) => {
    rl.question('Enter params (comma separated): ', (paramsInput) => {
        let params: any[];
        let param_types: string[];
        
        if (method === 'sort') {
            params = [paramsInput.split(',').map(p => p.trim())];
            param_types = ["string[]"];
        } else if (method === 'floor') {
            params = [parseFloat(paramsInput.trim())];
            param_types = ["float"];
        } else if (method === 'nroot') {
            const nums = paramsInput.split(',').map(p => parseFloat(p.trim()));
            params = nums;
            param_types = ["int", "int"];
        } else if (method === 'reverse' || method === 'validAnagram') {
            params = paramsInput.split(',').map(p => p.trim());
            param_types = params.map(() => "string");
        } else {
            params = paramsInput.split(',').map(p => p.trim());
            param_types = params.map(() => "string");
        }

        const Id = Math.floor(Math.random() * 100)
        
        console.log('Sending:', { method, params, param_types });

        const client = new net.Socket()

        client.connect(8090, "127.0.0.1", () => {
            console.log("Connection Established")
            const request: RPCRequest = {
                method: method,
                params: params,
                param_types: param_types,
                id: Id
            }
            client.write(JSON.stringify(request))
        })

        client.on("data", (response) => {
            try {
                const parsedResponse = response.toString()
                console.log(parsedResponse)
            }
            catch(err) {
                console.log("Error: ", err)
            }
        })
    });
});
