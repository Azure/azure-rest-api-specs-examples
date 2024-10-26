const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a AutoUpgradeProfile
 *
 * @summary Create a AutoUpgradeProfile
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/examples/AutoUpgradeProfiles_CreateOrUpdate.json
 */
async function createAnAutoUpgradeProfile() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const fleetName = "fleet1";
  const autoUpgradeProfileName = "autoupgradeprofile1";
  const resource = { channel: "Stable" };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.autoUpgradeProfiles.beginCreateOrUpdateAndWait(
    resourceGroupName,
    fleetName,
    autoUpgradeProfileName,
    resource,
  );
  console.log(result);
}
