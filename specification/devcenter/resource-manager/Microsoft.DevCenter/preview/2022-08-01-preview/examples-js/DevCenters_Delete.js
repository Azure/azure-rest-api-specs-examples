const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a devcenter
 *
 * @summary Deletes a devcenter
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/DevCenters_Delete.json
 */
async function devCentersDelete() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devCenters.beginDeleteAndWait(resourceGroupName, devCenterName);
  console.log(result);
}

devCentersDelete().catch(console.error);
