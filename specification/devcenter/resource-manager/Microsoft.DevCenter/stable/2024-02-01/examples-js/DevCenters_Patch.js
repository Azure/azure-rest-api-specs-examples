const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Partially updates a devcenter.
 *
 * @summary Partially updates a devcenter.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/DevCenters_Patch.json
 */
async function devCentersUpdate() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const devCenterName = "Contoso";
  const body = { tags: { costCode: "12345" } };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devCenters.beginUpdateAndWait(resourceGroupName, devCenterName, body);
  console.log(result);
}
