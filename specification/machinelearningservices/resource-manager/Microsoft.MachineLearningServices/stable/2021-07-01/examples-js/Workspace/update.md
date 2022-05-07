Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearningservices_5.0.1/sdk/machinelearningservices/arm-machinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a machine learning workspace with the specified parameters.
 *
 * @summary Updates a machine learning workspace with the specified parameters.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/update.json
 */
async function updateWorkspace() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workspace-1234";
  const workspaceName = "testworkspace";
  const parameters = {
    description: "new description",
    friendlyName: "New friendly name",
    publicNetworkAccess: "Disabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.workspaces.update(resourceGroupName, workspaceName, parameters);
  console.log(result);
}

updateWorkspace().catch(console.error);
```
