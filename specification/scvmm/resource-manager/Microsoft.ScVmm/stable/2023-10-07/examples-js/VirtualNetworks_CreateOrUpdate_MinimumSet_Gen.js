const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Onboards the ScVmm virtual network as an Azure virtual network resource.
 *
 * @summary Onboards the ScVmm virtual network as an Azure virtual network resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualNetworks_CreateOrUpdate_MinimumSet_Gen.json
 */
async function virtualNetworksCreateOrUpdateMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "rgscvmm";
  const virtualNetworkName = "-";
  const resource = { extendedLocation: {}, location: "fky" };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const result = await client.virtualNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualNetworkName,
    resource,
  );
  console.log(result);
}
