const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the admin state of the Network Interface.
 *
 * @summary Update the admin state of the Network Interface.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkInterfaces_updateAdministrativeState_MaximumSet_Gen.json
 */
async function networkInterfacesUpdateAdministrativeStateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const networkDeviceName = "networkDeviceName";
  const networkInterfaceName = "networkInterfaceName";
  const body = { state: "Enable" };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkInterfaces.beginUpdateAdministrativeStateAndWait(
    resourceGroupName,
    networkDeviceName,
    networkInterfaceName,
    body
  );
  console.log(result);
}
