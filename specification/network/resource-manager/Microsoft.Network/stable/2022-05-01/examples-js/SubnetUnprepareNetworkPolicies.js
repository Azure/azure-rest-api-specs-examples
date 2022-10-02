const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Unprepares a subnet by removing network intent policies.
 *
 * @summary Unprepares a subnet by removing network intent policies.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/SubnetUnprepareNetworkPolicies.json
 */
async function unprepareNetworkPolicies() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkName = "test-vnet";
  const subnetName = "subnet1";
  const unprepareNetworkPoliciesRequestParameters = {
    serviceName: "Microsoft.Sql/managedInstances",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.subnets.beginUnprepareNetworkPoliciesAndWait(
    resourceGroupName,
    virtualNetworkName,
    subnetName,
    unprepareNetworkPoliciesRequestParameters
  );
  console.log(result);
}

unprepareNetworkPolicies().catch(console.error);
