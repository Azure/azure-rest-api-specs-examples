const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to update the recovery settings of an ASR replication protected item.
 *
 * @summary the operation to update the recovery settings of an ASR replication protected item.
 * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Update_A2ADualStack.json
 */
async function updatesTheRecoveryIpConfigsOfADualStackIPv4AndIPv6A2AProtectedItem() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "42195872-7e70-4f8a-837f-84b28ecbb78b";
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectedItems.update(
    "priyanprg",
    "priyanponeboxvault",
    "CentralUSCanSite",
    "CentralUSCancloud",
    "a2aDualStackVm",
    {
      properties: {
        selectedRecoveryAzureNetworkId:
          "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/virtualNetworks/dualstack-recovery-vnet",
        vmNics: [
          {
            nicId:
              "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Network/networkInterfaces/a2aDualStackVm-nic",
            selectionType: "SelectedByUser",
            ipConfigs: [
              {
                ipConfigName: "ipconfig1",
                isPrimary: true,
                isSeletedForFailover: true,
                recoverySubnetName: "dualstack-subnet",
                recoveryStaticIPAddress: "10.0.2.46",
              },
              {
                ipConfigName: "ipconfig-v6",
                isPrimary: false,
                isSeletedForFailover: true,
                recoverySubnetName: "dualstack-subnet",
                recoveryStaticIPAddress: "2001:db8:1::4",
              },
            ],
          },
        ],
        providerSpecificDetails: {
          instanceType: "A2A",
          recoveryResourceGroupId:
            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg",
        },
      },
    },
  );
  console.log(result);
}
