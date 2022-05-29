Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearning_1.0.0-beta.1/sdk/machinelearning/arm-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all skus with associated features
 *
 * @summary Lists all skus with associated features
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/WorkspaceSku/list.json
 */
async function listSkus() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaceSkus.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSkus().catch(console.error);
```
