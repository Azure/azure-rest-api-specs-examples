const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates policy configuration for the workspace.
 *
 * @summary creates or updates policy configuration for the workspace.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateWorkspacePolicy.json
 */
async function apiManagementCreateWorkspacePolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspacePolicy.createOrUpdate(
    "rg1",
    "apimService1",
    "wks1",
    "policy",
    {
      format: "xml",
      value:
        "<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>",
    },
    { ifMatch: "*" },
  );
  console.log(result);
}
