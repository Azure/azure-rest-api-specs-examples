const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the specified Bastion Host.
 *
 * @summary creates or updates the specified Bastion Host.
 * x-ms-original-file: 2026-01-01/BastionHostPutWithUserAssignedIdentityForSRConfig.json
 */
async function createOrUpdateBastionHostWithUserAssignedIdentityForSessionRecordingConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.bastionHosts.createOrUpdate("rg1", "bastionhosttenant", {
    enableSessionRecording: true,
    sessionRecordingConfiguration: {
      identity: {
        type: "UserAssigned",
        userAssignedIdentityId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userassignedmsi",
      },
      blobContainerUri: "https://contosostorage.blob.core.windows.net/contosocontainer",
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
