const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a devcenter resource
 *
 * @summary Creates or updates a devcenter resource
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/DevCenters_Create.json
 */
async function devCentersCreate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const body = {
    location: "centralus",
    tags: { costCode: "12345" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devCenters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    devCenterName,
    body
  );
  console.log(result);
}

devCentersCreate().catch(console.error);
