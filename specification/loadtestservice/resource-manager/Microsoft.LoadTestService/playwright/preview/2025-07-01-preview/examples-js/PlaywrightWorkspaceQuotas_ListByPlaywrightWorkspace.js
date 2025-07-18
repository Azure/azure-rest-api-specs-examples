const { PlaywrightManagementClient } = require("@azure/arm-playwright");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list quota resources for a given Playwright workspace resource.
 *
 * @summary list quota resources for a given Playwright workspace resource.
 * x-ms-original-file: 2025-07-01-preview/PlaywrightWorkspaceQuotas_ListByPlaywrightWorkspace.json
 */
async function playwrightWorkspaceQuotasListByPlaywrightWorkspace() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new PlaywrightManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.playwrightWorkspaceQuotas.listByPlaywrightWorkspace(
    "dummyrg",
    "myWorkspace",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
