const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get Site Analyses
 *
 * @summary Description for Get Site Analyses
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_ListSiteAnalysesSlot.json
 */
async function listAppSlotAnalyses() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "Sample-WestUSResourceGroup";
  const siteName = "SampleApp";
  const diagnosticCategory = "availability";
  const slot = "staging";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diagnostics.listSiteAnalysesSlot(
    resourceGroupName,
    siteName,
    diagnosticCategory,
    slot
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAppSlotAnalyses().catch(console.error);
