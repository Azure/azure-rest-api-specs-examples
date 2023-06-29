const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the L3 Isolation Domain resource.
 *
 * @summary API to update certain properties of the L3 Isolation Domain resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L3IsolationDomains_Update_MaximumSet_Gen.json
 */
async function l3IsolationDomainsUpdateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const l3IsolationDomainName = "example-l3domain";
  const body = {
    description: "creating L3 isolation domain",
    aggregateRouteConfiguration: {
      ipv4Routes: [{ prefix: "10.0.0.0/24" }],
      ipv6Routes: [{ prefix: "3FFE:FFFF:0:CD30::a0/29" }],
    },
    connectedSubnetRoutePolicy: {
      exportRoutePolicyId:
        "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
    },
    redistributeConnectedSubnets: "True",
    redistributeStaticRoutes: "False",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.l3IsolationDomains.beginUpdateAndWait(
    resourceGroupName,
    l3IsolationDomainName,
    body
  );
  console.log(result);
}
