const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Bastion Host.
 *
 * @summary creates or updates the specified Bastion Host.
 * x-ms-original-file: 2026-01-01/BastionHostPutWithSystemAssignedIdentityForSRConfig.json
 */
async function createOrUpdateBastionHostWithSystemAssignedIdentityForSessionRecordingConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.bastionHosts.createOrUpdate("rg1", "bastionhosttenant", {
    enableSessionRecording: true,
    sessionRecordingConfiguration: {
      identity: { type: "SystemAssigned" },
      blobContainerUri: "https://mystorageaccount.blob.core.windows.net/mycontainer",
    },
    ipConfigurations: [
      {
        name: "bastionHostIpConfiguration",
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet",
        },
        publicIPAddress: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName",
        },
      },
    ],
  });
  console.log(result);
}
