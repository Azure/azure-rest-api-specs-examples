const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private link service.
 *
 * @summary Deletes the specified private link service.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/PrivateLinkServiceDelete.json
 */
async function deletePrivateLinkService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const serviceName = "testPls";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateLinkServices.beginDeleteAndWait(
    resourceGroupName,
    serviceName
  );
  console.log(result);
}

deletePrivateLinkService().catch(console.error);
