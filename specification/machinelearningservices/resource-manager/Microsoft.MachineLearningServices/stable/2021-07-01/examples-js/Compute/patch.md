Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearning_1.0.0-beta.1/sdk/machinelearning/arm-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation.
 *
 * @summary Updates properties of a compute. This call will overwrite a compute if it exists. This is a nonrecoverable operation.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Compute/patch.json
 */
async function updateAAmlCompute() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const workspaceName = "workspaces123";
  const computeName = "compute123";
  const parameters = {
    properties: {
      scaleSettings: {
        maxNodeCount: 4,
        minNodeCount: 4,
        nodeIdleTimeBeforeScaleDown: "PT5M",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.computeOperations.beginUpdateAndWait(
    resourceGroupName,
    workspaceName,
    computeName,
    parameters
  );
  console.log(result);
}

updateAAmlCompute().catch(console.error);
```
