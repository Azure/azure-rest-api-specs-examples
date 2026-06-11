const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or Updates a group.
 *
 * @summary creates or Updates a group.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateWorkspaceGroup.json
 */
async function apiManagementCreateWorkspaceGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceGroup.createOrUpdate(
    "rg1",
    "apimService1",
    "wks1",
    "tempgroup",
    { displayName: "temp group" },
  );
  console.log(result);
}
