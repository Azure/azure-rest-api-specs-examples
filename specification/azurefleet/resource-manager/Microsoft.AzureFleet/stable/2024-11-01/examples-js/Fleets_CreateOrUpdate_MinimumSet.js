const { AzureFleetClient } = require("@azure/arm-computefleet");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Fleet
 *
 * @summary create a Fleet
 * x-ms-original-file: 2024-11-01/Fleets_CreateOrUpdate_MinimumSet.json
 */
async function fleetsCreateOrUpdateMinimumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1DC2F28C-A625-4B0E-9748-9885A3C9E9EB";
  const client = new AzureFleetClient(credential, subscriptionId);
  const result = await client.fleets.createOrUpdate("rgazurefleet", "testFleet", {
    properties: {
      spotPriorityProfile: {
        capacity: 2,
        minCapacity: 1,
        evictionPolicy: "Delete",
        allocationStrategy: "PriceCapacityOptimized",
        maintain: true,
      },
      regularPriorityProfile: {
        capacity: 2,
        minCapacity: 1,
        allocationStrategy: "LowestPrice",
      },
      vmSizesProfile: [
        { name: "Standard_D2s_v3" },
        { name: "Standard_D4s_v3" },
        { name: "Standard_E2s_v3" },
      ],
      computeProfile: {
        baseVirtualMachineProfile: {
          storageProfile: {
            imageReference: {
              publisher: "canonical",
              offer: "0001-com-ubuntu-server-focal",
              sku: "20_04-lts-gen2",
              version: "latest",
            },
            osDisk: {
              caching: "ReadWrite",
              createOption: "FromImage",
              osType: "Linux",
              managedDisk: { storageAccountType: "Standard_LRS" },
            },
          },
          osProfile: {
            computerNamePrefix: "prefix",
            adminUsername: "azureuser",
            adminPassword: "TestPassword$0",
            linuxConfiguration: { disablePasswordAuthentication: false },
          },
          networkProfile: {
            networkInterfaceConfigurations: [
              {
                name: "vmNameTest",
                properties: {
                  primary: true,
                  enableAcceleratedNetworking: false,
                  ipConfigurations: [
                    {
                      name: "vmNameTest",
                      properties: {
                        subnet: {
                          id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}",
                        },
                        primary: true,
                        loadBalancerBackendAddressPools: [
                          {
                            id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}",
                          },
                        ],
                      },
                    },
                  ],
                  enableIPForwarding: true,
                },
              },
            ],
          },
        },
        computeApiVersion: "2023-09-01",
        platformFaultDomainCount: 1,
      },
    },
    tags: { key: "fleets-test" },
    location: "eastus2euap",
  });
  console.log(result);
}
