Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicelinker_2.0.0/sdk/servicelinker/arm-servicelinker/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available ServiceLinker REST API operations.
 *
 * @summary Lists the available ServiceLinker REST API operations.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/OperationsList.json
 */
async function getConfiguration() {
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getConfiguration().catch(console.error);
```
