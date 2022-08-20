const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Network Connections resource
 *
 * @summary Creates or updates a Network Connections resource
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/NetworkConnections_Put.json
 */
async function networkConnectionsCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const networkConnectionName = "uswest3network";
  const body = {
    domainJoinType: "HybridAzureADJoin",
    domainName: "mydomaincontroller.local",
    domainPassword: "Password value for user",
    domainUsername: "testuser@mydomaincontroller.local",
    location: "centralus",
    networkingResourceGroupName: "NetworkInterfaces",
    subnetId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ExampleRG/providers/Microsoft.Network/virtualNetworks/ExampleVNet/subnets/default",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.networkConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkConnectionName,
    body
  );
  console.log(result);
}

networkConnectionsCreateOrUpdate().catch(console.error);
