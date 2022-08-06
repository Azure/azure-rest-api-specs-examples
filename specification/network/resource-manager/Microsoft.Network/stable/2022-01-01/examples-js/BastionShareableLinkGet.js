const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return the Bastion Shareable Links for all the VMs specified in the request.
 *
 * @summary Return the Bastion Shareable Links for all the VMs specified in the request.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/BastionShareableLinkGet.json
 */
async function returnsTheBastionShareableLinksForTheRequestVMS() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const bastionHostName = "bastionhosttenant";
  const bslRequest = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.listBastionShareableLink(
    resourceGroupName,
    bastionHostName,
    bslRequest
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

returnsTheBastionShareableLinksForTheRequestVMS().catch(console.error);
