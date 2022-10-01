const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Bastion Shareable Links for all the VMs specified in the request.
 *
 * @summary Creates a Bastion Shareable Links for all the VMs specified in the request.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/BastionShareableLinkCreate.json
 */
async function createBastionShareableLinksForTheRequestVMS() {
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
  for await (let item of client.beginListPutBastionShareableLinkAndWait(
    resourceGroupName,
    bastionHostName,
    bslRequest
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

createBastionShareableLinksForTheRequestVMS().catch(console.error);
