Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearningservices_5.0.1/sdk/machinelearningservices/arm-machinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearningservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all enabled features for a workspace
 *
 * @summary Lists all enabled features for a workspace
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/WorkspaceFeature/list.json
 */
async function listWorkspaceFeatures() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const workspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaceFeatures.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listWorkspaceFeatures().catch(console.error);
```
