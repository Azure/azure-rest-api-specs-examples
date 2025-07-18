const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the base Settings of the target resource.
 *
 * @summary Updates the base Settings of the target resource.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/settings/SettingsUpdate.json
 */
async function settingsUpdate() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "hybridRG";
  const baseProvider = "Microsoft.HybridCompute";
  const baseResourceType = "machines";
  const baseResourceName = "testMachine";
  const settingsResourceName = "default";
  const parameters = {
    gatewayResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/gateways/newGateway",
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.settingsOperations.update(
    resourceGroupName,
    baseProvider,
    baseResourceType,
    baseResourceName,
    settingsResourceName,
    parameters,
  );
  console.log(result);
}
