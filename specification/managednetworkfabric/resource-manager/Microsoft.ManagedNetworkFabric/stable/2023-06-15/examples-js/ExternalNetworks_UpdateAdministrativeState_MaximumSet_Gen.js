const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Executes update operation to enable or disable administrative State for externalNetwork.
 *
 * @summary Executes update operation to enable or disable administrative State for externalNetwork.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ExternalNetworks_UpdateAdministrativeState_MaximumSet_Gen.json
 */
async function externalNetworksUpdateAdministrativeStateMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const l3IsolationDomainName = "example-l3domain";
  const externalNetworkName = "example-externalnetwork";
  const body = {
    resourceIds: [""],
    state: "Enable",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.externalNetworks.beginUpdateAdministrativeStateAndWait(
    resourceGroupName,
    l3IsolationDomainName,
    externalNetworkName,
    body
  );
  console.log(result);
}
