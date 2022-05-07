Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-operationalinsights_8.0.1/sdk/operationalinsights/arm-operationalinsights/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the status of a long running azure asynchronous operation.
 *
 * @summary Get the status of a long running azure asynchronous operation.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/OperationStatusesGet.json
 */
async function getSpecificOperationStatus() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const location = "West US";
  const asyncOperationId = "713192d7-503f-477a-9cfe-4efc3ee2bd11";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.operationStatuses.get(location, asyncOperationId);
  console.log(result);
}

getSpecificOperationStatus().catch(console.error);
```
