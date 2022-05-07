Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementgroups_2.0.1/sdk/managementgroups/arm-managementgroups/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a management group.
If a management group is already created and a subsequent create request is issued with different properties, the management group properties will be updated.

 *
 * @summary Create or update a management group.
If a management group is already created and a subsequent create request is issued with different properties, the management group properties will be updated.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutManagementGroup.json
 */
async function putManagementGroup() {
  const groupId = "ChildGroup";
  const cacheControl = "no-cache";
  const createManagementGroupRequest = {
    displayName: "ChildGroup",
    details: {
      parent: {
        id: "/providers/Microsoft.Management/managementGroups/RootGroup",
      },
    },
  };
  const options = {
    cacheControl,
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.managementGroups.beginCreateOrUpdateAndWait(
    groupId,
    createManagementGroupRequest,
    options
  );
  console.log(result);
}

putManagementGroup().catch(console.error);
```
