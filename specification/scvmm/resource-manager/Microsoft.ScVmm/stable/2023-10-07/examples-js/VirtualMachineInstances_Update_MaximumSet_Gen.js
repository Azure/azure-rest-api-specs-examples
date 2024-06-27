const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a virtual machine instance.
 *
 * @summary The operation to update a virtual machine instance.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_Update_MaximumSet_Gen.json
 */
async function virtualMachineInstancesUpdateMaximumSet() {
  const resourceUri = "gtgclehcbsyave";
  const properties = {
    properties: {
      availabilitySets: [
        {
          name: "lwbhaseo",
          id: "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/availabilitySets/availabilitySetResourceName",
        },
      ],
      hardwareProfile: {
        cpuCount: 22,
        dynamicMemoryEnabled: "true",
        dynamicMemoryMaxMB: 2,
        dynamicMemoryMinMB: 30,
        limitCpuForMigration: "true",
        memoryMB: 5,
      },
      infrastructureProfile: { checkpointType: "jkbpzjxpeegackhsvikrnlnwqz" },
      networkProfile: {
        networkInterfaces: [
          {
            name: "kvofzqulbjlbtt",
            ipv4AddressType: "Dynamic",
            ipv6AddressType: "Dynamic",
            macAddress: "oaeqqegt",
            macAddressType: "Dynamic",
            nicId: "roxpsvlo",
            virtualNetworkId:
              "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName",
          },
        ],
      },
      storageProfile: {
        disks: [
          {
            name: "fgnckfymwdsqnfxkdvexuaobe",
            bus: 8,
            busType: "zu",
            diskId: "ltdrwcfjklpsimhzqyh",
            diskSizeGB: 30,
            lun: 10,
            storageQosPolicy: { name: "ceiyfrflu", id: "o" },
            vhdType: "cnbeeeylrvopigdynvgpkfp",
          },
        ],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const result = await client.virtualMachineInstances.beginUpdateAndWait(resourceUri, properties);
  console.log(result);
}
