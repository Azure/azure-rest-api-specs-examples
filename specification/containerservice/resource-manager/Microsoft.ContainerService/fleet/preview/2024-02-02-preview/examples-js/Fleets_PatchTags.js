const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Fleet
 *
 * @summary Update a Fleet
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-02-02-preview/examples/Fleets_PatchTags.json
 */
async function updateAFleet() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const ifMatch = "dfjkwelr7384";
  const fleetName = "fleet1";
  const properties = { tags: { env: "prod", tier: "secure" } };
  const options = { ifMatch };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleets.beginUpdateAndWait(
    resourceGroupName,
    fleetName,
    properties,
    options,
  );
  console.log(result);
}
