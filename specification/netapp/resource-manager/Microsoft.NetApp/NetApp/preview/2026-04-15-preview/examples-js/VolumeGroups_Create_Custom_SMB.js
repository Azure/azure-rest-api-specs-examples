const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a volume group along with specified volumes
 *
 * @summary create a volume group along with specified volumes
 * x-ms-original-file: 2026-04-15-preview/VolumeGroups_Create_Custom_SMB.json
 */
async function volumeGroupsCreateCustomSMB() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumeGroups.create("myRG", "account1", "group1", {
    location: "westus",
    properties: {
      groupMetaData: {
        applicationIdentifier: "CU2",
        applicationType: "CUSTOM",
        groupDescription: "Volume group",
      },
      volumes: [
        {
          name: "test-cus-data1",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data1",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data1",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data2",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data2",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data2",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data3",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data3",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data3",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data4",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data4",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data4",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data5",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data5",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data5",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data6",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data6",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data6",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data7",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data7",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data7",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data8",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data8",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data8",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data9",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data9",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data9",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data10",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data10",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data10",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data11",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data11",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data11",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
        {
          name: "test-cus-data12",
          properties: {
            capacityPoolResourceId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1",
            creationToken: "test-cus-data12",
            protocolTypes: ["CIFS"],
            serviceLevel: "Premium",
            subnetId:
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3",
            throughputMibps: 10,
            usageThreshold: 107374182400,
            volumeSpecName: "cus-data12",
            smbEncryption: false,
            smbContinuouslyAvailable: false,
            smbNonBrowsable: "Disabled",
            smbAccessBasedEnumeration: "Disabled",
            avsDataStore: "Disabled",
          },
          zones: ["1"],
        },
      ],
    },
  });
  console.log(result);
}
