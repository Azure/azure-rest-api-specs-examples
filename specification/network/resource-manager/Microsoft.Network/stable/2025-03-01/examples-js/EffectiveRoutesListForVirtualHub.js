const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the effective routes configured for the Virtual Hub resource or the specified resource .
 *
 * @summary Gets the effective routes configured for the Virtual Hub resource or the specified resource .
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/EffectiveRoutesListForVirtualHub.json
 */
async function effectiveRoutesForTheVirtualHub() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualHubName = "virtualHub1";
  const effectiveRoutesParameters = {};
  const options = {
    effectiveRoutesParameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubs.beginGetEffectiveVirtualHubRoutesAndWait(
    resourceGroupName,
    virtualHubName,
    options,
  );
  console.log(result);
}
