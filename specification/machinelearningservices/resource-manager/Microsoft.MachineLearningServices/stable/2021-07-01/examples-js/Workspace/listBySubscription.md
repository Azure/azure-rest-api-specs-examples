Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearningservices_5.0.1/sdk/machinelearningservices/arm-machinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available machine learning workspaces under the specified subscription.
 *
 * @summary Lists all the available machine learning workspaces under the specified subscription.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/listBySubscription.json
 */
async function getWorkspacesBySubscription() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getWorkspacesBySubscription().catch(console.error);
```
