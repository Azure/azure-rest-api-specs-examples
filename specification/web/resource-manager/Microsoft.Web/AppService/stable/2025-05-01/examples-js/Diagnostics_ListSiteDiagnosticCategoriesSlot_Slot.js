const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to description for Get Diagnostics Categories
 *
 * @summary description for Get Diagnostics Categories
 * x-ms-original-file: 2025-05-01/Diagnostics_ListSiteDiagnosticCategoriesSlot_Slot.json
 */
async function listAppSlotDiagnosticCategories() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.diagnostics.listSiteDiagnosticCategoriesSlot(
    "Sample-WestUSResourceGroup",
    "SampleApp",
    "staging",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
