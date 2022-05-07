Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementgroups_2.0.1/sdk/managementgroups/arm-managementgroups/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a management group.

 *
 * @summary Update a management group.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PatchManagementGroup.json
 */
async function patchManagementGroup() {
  const groupId = "ChildGroup";
  const cacheControl = "no-cache";
  const patchGroupRequest = {
    displayName: "AlternateDisplayName",
    parentGroupId: "/providers/Microsoft.Management/managementGroups/AlternateRootGroup",
  };
  const options = { cacheControl };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.managementGroups.update(groupId, patchGroupRequest, options);
  console.log(result);
}

patchManagementGroup().catch(console.error);
```
