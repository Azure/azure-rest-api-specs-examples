const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes the Bastion Shareable Links for all the VMs specified in the request.
 *
 * @summary Deletes the Bastion Shareable Links for all the VMs specified in the request.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/BastionShareableLinkDelete.json
 */
async function deleteBastionShareableLinksForTheRequestVMS() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const bastionHostName = "bastionhosttenant";
  const options = {
    body: {
      vms: [
        {
          vm: {
            id: "/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1",
          },
        },
        {
          vm: {
            id: "/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2",
          },
        },
      ],
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/deleteShareableLinks",
      subscriptionId,
      resourceGroupName,
      bastionHostName,
    )
    .post(options);
  const poller = await getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

deleteBastionShareableLinksForTheRequestVMS().catch(console.error);
