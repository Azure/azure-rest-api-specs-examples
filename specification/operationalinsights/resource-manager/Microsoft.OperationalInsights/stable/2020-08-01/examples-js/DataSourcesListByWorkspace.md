Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-operationalinsights_8.0.1/sdk/operationalinsights/arm-operationalinsights/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the first page of data source instances in a workspace with the link to the next page.
 *
 * @summary Gets the first page of data source instances in a workspace with the link to the next page.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataSourcesListByWorkspace.json
 */
async function dataSourcesListByWorkspace() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "OIAutoRest5123";
  const workspaceName = "AzTest9724";
  const filter = "kind='WindowsEvent'";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataSources.listByWorkspace(
    resourceGroupName,
    workspaceName,
    filter
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

dataSourcesListByWorkspace().catch(console.error);
```
