const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the provided bare metal machine.
All customer initiated requests will be rejected as the life cycle of this resource is managed by the system.
 *
 * @summary Delete the provided bare metal machine.
All customer initiated requests will be rejected as the life cycle of this resource is managed by the system.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BareMetalMachines_Delete.json
 */
async function deleteBareMetalMachine() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const bareMetalMachineName = "bareMetalMachineName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.bareMetalMachines.beginDeleteAndWait(
    resourceGroupName,
    bareMetalMachineName
  );
  console.log(result);
}
