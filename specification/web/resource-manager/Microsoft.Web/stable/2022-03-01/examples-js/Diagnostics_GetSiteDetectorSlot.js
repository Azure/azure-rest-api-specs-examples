const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get Detector
 *
 * @summary Description for Get Detector
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/Diagnostics_GetSiteDetectorSlot.json
 */
async function getAppSlotDetector() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "Sample-WestUSResourceGroup";
  const siteName = "SampleApp";
  const diagnosticCategory = "availability";
  const detectorName = "sitecrashes";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.diagnostics.getSiteDetector(
    resourceGroupName,
    siteName,
    diagnosticCategory,
    detectorName
  );
  console.log(result);
}

getAppSlotDetector().catch(console.error);
