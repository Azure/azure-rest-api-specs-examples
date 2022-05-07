Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementgroups_2.0.1/sdk/managementgroups/arm-managementgroups/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves details about all subscriptions which are associated with the management group.

 *
 * @summary Retrieves details about all subscriptions which are associated with the management group.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetAllSubscriptionsFromManagementGroup.json
 */
async function getAllSubscriptionsFromManagementGroup() {
  const groupId = "Group";
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const resArray = new Array();
  for await (let item of client.managementGroupSubscriptions.listSubscriptionsUnderManagementGroup(
    groupId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllSubscriptionsFromManagementGroup().catch(console.error);
```
