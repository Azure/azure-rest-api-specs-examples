```javascript
const { ChangesClient } = require("@azure/arm-changes");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Obtains the specified change resource for the target resource
 *
 * @summary Obtains the specified change resource for the target resource
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-05-01/examples/GetChange.json
 */
async function getChange() {
  const subscriptionId = "subscriptionId1";
  const resourceGroupName = "resourceGroup1";
  const resourceProviderNamespace = "resourceProvider1";
  const resourceType = "resourceType1";
  const resourceName = "resourceName1";
  const changeResourceId = "1d58d72f-0719-4a48-9228-b7ea682885bf";
  const credential = new DefaultAzureCredential();
  const client = new ChangesClient(credential, subscriptionId);
  const result = await client.changes.get(
    resourceGroupName,
    resourceProviderNamespace,
    resourceType,
    resourceName,
    changeResourceId
  );
  console.log(result);
}

getChange().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-changes_1.0.0/sdk/changes/arm-changes/README.md) on how to add the SDK to your project and authenticate.
