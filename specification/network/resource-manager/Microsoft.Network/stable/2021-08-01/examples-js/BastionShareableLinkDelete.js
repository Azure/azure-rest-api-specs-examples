const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the Bastion Shareable Links for all the VMs specified in the request.
 *
 * @summary Deletes the Bastion Shareable Links for all the VMs specified in the request.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/BastionShareableLinkDelete.json
 */
async function deleteBastionShareableLinksForTheRequestVMS() {
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
  const result = await client.beginDeleteBastionShareableLinkAndWait(
    resourceGroupName,
    bastionHostName,
    bslRequest
  );
  console.log(result);
}

deleteBastionShareableLinksForTheRequestVMS().catch(console.error);
