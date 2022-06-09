```javascript
const { ChangesClient } = require("@azure/arm-changes");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Obtains a list of change resources from the past 14 days for the target resource
 *
 * @summary Obtains a list of change resources from the past 14 days for the target resource
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-05-01/examples/ListChanges.json
 */
async function listChanges() {
  const subscriptionId = "subscriptionId1";
  const resourceGroupName = "resourceGroup1";
  const resourceProviderNamespace = "resourceProvider1";
  const resourceType = "resourceType1";
  const resourceName = "resourceName1";
  const credential = new DefaultAzureCredential();
  const client = new ChangesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.changes.list(
    resourceGroupName,
    resourceProviderNamespace,
    resourceType,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listChanges().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-changes_1.0.0/sdk/changes/arm-changes/README.md) on how to add the SDK to your project and authenticate.
