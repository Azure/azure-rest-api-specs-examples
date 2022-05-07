Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearningservices_5.0.1/sdk/machinelearningservices/arm-machinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Posts a start action to a compute instance
 *
 * @summary Posts a start action to a compute instance
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Compute/start.json
 */
async function startComputeInstanceCompute() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspaces123";
  const computeName = "compute123";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.computeOperations.beginStartAndWait(
    resourceGroupName,
    workspaceName,
    computeName
  );
  console.log(result);
}

startComputeInstanceCompute().catch(console.error);
```
