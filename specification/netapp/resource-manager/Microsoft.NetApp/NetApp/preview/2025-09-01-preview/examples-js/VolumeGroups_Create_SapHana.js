const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a volume group along with specified volumes
 *
 * @summary create a volume group along with specified volumes
 * x-ms-original-file: 2025-09-01-preview/VolumeGroups_Create_SapHana.json
 */
async function volumeGroupsCreateSapHana() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumeGroups.create("myRG", "account1", "group1", {
    location: "westus",
    properties: {
      groupMetaData: {
        applicationIdentifier: "SH9",
        applicationType: "SAP-HANA",
        groupDescription: "Volume group",
      },
      volumes: [
        {
          name: "test-data-mnt00001",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-data-mnt00001",
            exportPolicy: {
              rules: [
                {
                  allowedClients: "0.0.0.0/0",
                  cifs: false,
                  hasRootAccess: true,
                  kerberos5ReadOnly: false,
                  kerberos5ReadWrite: false,
                  kerberos5IReadOnly: false,
                  kerberos5IReadWrite: false,
                  kerberos5PReadOnly: false,
                  kerberos5PReadWrite: false,
                  nfsv3: false,
                  nfsv41: true,
                  ruleIndex: 1,
                  unixReadOnly: true,
                  unixReadWrite: true,
                },
              ],
            },
            protocolTypes: ["NFSv4.1"],
            proximityPlacementGroup:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "data",
          },
        },
        {
          name: "test-log-mnt00001",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-log-mnt00001",
            exportPolicy: {
              rules: [
                {
                  allowedClients: "0.0.0.0/0",
                  cifs: false,
                  hasRootAccess: true,
                  kerberos5ReadOnly: false,
                  kerberos5ReadWrite: false,
                  kerberos5IReadOnly: false,
                  kerberos5IReadWrite: false,
                  kerberos5PReadOnly: false,
                  kerberos5PReadWrite: false,
                  nfsv3: false,
                  nfsv41: true,
                  ruleIndex: 1,
                  unixReadOnly: true,
                  unixReadWrite: true,
                },
              ],
            },
            protocolTypes: ["NFSv4.1"],
            proximityPlacementGroup:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "log",
          },
        },
        {
          name: "test-shared",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-shared",
            exportPolicy: {
              rules: [
                {
                  allowedClients: "0.0.0.0/0",
                  cifs: false,
                  hasRootAccess: true,
                  kerberos5ReadOnly: false,
                  kerberos5ReadWrite: false,
                  kerberos5IReadOnly: false,
                  kerberos5IReadWrite: false,
                  kerberos5PReadOnly: false,
                  kerberos5PReadWrite: false,
                  nfsv3: false,
                  nfsv41: true,
                  ruleIndex: 1,
                  unixReadOnly: true,
                  unixReadWrite: true,
                },
              ],
            },
            protocolTypes: ["NFSv4.1"],
            proximityPlacementGroup:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "shared",
          },
        },
        {
          name: "test-data-backup",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-data-backup",
            exportPolicy: {
              rules: [
                {
                  allowedClients: "0.0.0.0/0",
                  cifs: false,
                  hasRootAccess: true,
                  kerberos5ReadOnly: false,
                  kerberos5ReadWrite: false,
                  kerberos5IReadOnly: false,
                  kerberos5IReadWrite: false,
                  kerberos5PReadOnly: false,
                  kerberos5PReadWrite: false,
                  nfsv3: false,
                  nfsv41: true,
                  ruleIndex: 1,
                  unixReadOnly: true,
                  unixReadWrite: true,
                },
              ],
            },
            protocolTypes: ["NFSv4.1"],
            proximityPlacementGroup:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "data-backup",
          },
        },
        {
          name: "test-log-backup",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-log-backup",
            exportPolicy: {
              rules: [
                {
                  allowedClients: "0.0.0.0/0",
                  cifs: false,
                  hasRootAccess: true,
                  kerberos5ReadOnly: false,
                  kerberos5ReadWrite: false,
                  kerberos5IReadOnly: false,
                  kerberos5IReadWrite: false,
                  kerberos5PReadOnly: false,
                  kerberos5PReadWrite: false,
                  nfsv3: false,
                  nfsv41: true,
                  ruleIndex: 1,
                  unixReadOnly: true,
                  unixReadWrite: true,
                },
              ],
            },
            protocolTypes: ["NFSv4.1"],
            proximityPlacementGroup:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg",
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "log-backup",
          },
        },
      ],
    },
  });
  console.log(result);
}
