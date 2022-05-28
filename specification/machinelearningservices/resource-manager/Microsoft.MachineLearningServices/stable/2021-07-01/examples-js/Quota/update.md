Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-machinelearning_1.0.0-beta.1/sdk/machinelearning/arm-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update quota for each VM family in workspace.
 *
 * @summary Update quota for each VM family in workspace.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Quota/update.json
 */
async function updateQuotas() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const parameters = {
    value: [
      {
        type: "Microsoft.MachineLearningServices/workspaces/quotas",
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs",
        limit: 100,
        unit: "Count",
      },
      {
        type: "Microsoft.MachineLearningServices/workspaces/quotas",
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs",
        limit: 200,
        unit: "Count",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.quotas.update(location, parameters);
  console.log(result);
}

updateQuotas().catch(console.error);
```
