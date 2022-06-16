const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get site detector response
 *
 * @summary Description for Get site detector response
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/Diagnostics_GetSiteDetectorResponseSlot.json
 */
async function getAppSlotDetectorResponse() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "Sample-WestUSResourceGroup";
  const siteName = "SampleApp";
  const detectorName = "runtimeavailability";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.diagnostics.getSiteDetectorResponse(
    resourceGroupName,
    siteName,
    detectorName
  );
  console.log(result);
}

getAppSlotDetectorResponse().catch(console.error);
