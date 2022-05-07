Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementgroups_2.0.1/sdk/managementgroups/arm-managementgroups/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Associates existing subscription with the management group.

 *
 * @summary Associates existing subscription with the management group.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/AddManagementGroupSubscription.json
 */
async function addSubscriptionToManagementGroup() {
  const groupId = "Group";
  const subscriptionId = "728bcbe4-8d56-4510-86c2-4921b8beefbc";
  const cacheControl = "no-cache";
  const options = {
    cacheControl,
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.managementGroupSubscriptions.create(groupId, subscriptionId, options);
  console.log(result);
}

addSubscriptionToManagementGroup().catch(console.error);
```
