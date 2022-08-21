const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Fleet member.
 *
 * @summary Gets a Fleet member.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/FleetMembers_Get.json
 */
async function getsAFleetMemberResource() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const fleetName = "fleet-1";
  const fleetMemberName = "member-1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.fleetMembers.get(resourceGroupName, fleetName, fleetMemberName);
  console.log(result);
}

getsAFleetMemberResource().catch(console.error);
