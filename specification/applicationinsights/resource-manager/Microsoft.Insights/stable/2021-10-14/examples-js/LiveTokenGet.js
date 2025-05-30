const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to **Gets an access token for live metrics stream data.**
 *
 * @summary **Gets an access token for live metrics stream data.**
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-10-14/examples/LiveTokenGet.json
 */
async function getLiveTokenForResource() {
  const resourceUri =
    "subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/FabrikamFiberApp/providers/microsoft.insights/components/CustomAvailabilityTest/providers/microsoft.insights/generatelivetoken";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential);
  const result = await client.liveToken.get(resourceUri);
  console.log(result);
}
