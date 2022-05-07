Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managedapplications_2.0.1/sdk/managedapplications/arm-managedapplications/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the managed application.
 *
 * @summary Deletes the managed application.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/deleteApplicationById.json
 */
async function deleteApplicationById() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const applicationId = "myApplicationId";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.beginDeleteByIdAndWait(applicationId);
  console.log(result);
}

deleteApplicationById().catch(console.error);
```
