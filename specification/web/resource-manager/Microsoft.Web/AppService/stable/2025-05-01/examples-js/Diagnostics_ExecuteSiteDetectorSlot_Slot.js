const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to description for Execute Detector
 *
 * @summary description for Execute Detector
 * x-ms-original-file: 2025-05-01/Diagnostics_ExecuteSiteDetectorSlot_Slot.json
 */
async function executeSiteSlotDetector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.diagnostics.executeSiteDetectorSlot(
    "Sample-WestUSResourceGroup",
    "SampleApp",
    "sitecrashes",
    "availability",
    "staging",
  );
  console.log(result);
}
