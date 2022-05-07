Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementgroups_2.0.1/sdk/managementgroups/arm-managementgroups/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the hierarchy settings defined at the Management Group level.

 *
 * @summary Creates or updates the hierarchy settings defined at the Management Group level.

 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutHierarchySettings.json
 */
async function getGroupSettings() {
  const groupId = "root";
  const createTenantSettingsRequest = {
    defaultManagementGroup: "/providers/Microsoft.Management/managementGroups/DefaultGroup",
    requireAuthorizationForGroupCreation: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.hierarchySettingsOperations.createOrUpdate(
    groupId,
    createTenantSettingsRequest
  );
  console.log(result);
}

getGroupSettings().catch(console.error);
```
