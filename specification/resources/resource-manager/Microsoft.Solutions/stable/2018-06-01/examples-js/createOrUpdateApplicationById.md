Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managedapplications_2.0.1/sdk/managedapplications/arm-managedapplications/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new managed application.
 *
 * @summary Creates a new managed application.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplicationById.json
 */
async function createOrUpdateApplicationById() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const applicationId = "myApplicationId";
  const parameters = {
    kind: "ServiceCatalog",
    location: "East US 2",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.beginCreateOrUpdateByIdAndWait(
    applicationId,
    parameters
  );
  console.log(result);
}

createOrUpdateApplicationById().catch(console.error);
```
