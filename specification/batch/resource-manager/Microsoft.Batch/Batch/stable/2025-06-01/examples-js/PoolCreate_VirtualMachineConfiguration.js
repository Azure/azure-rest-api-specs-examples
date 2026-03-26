const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_VirtualMachineConfiguration.json
 */
async function createPoolFullVirtualMachineConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          dataDisks: [
            {
              caching: "ReadWrite",
              diskSizeGB: 30,
              lun: 0,
              managedDisk: { storageAccountType: "StandardSSD_LRS" },
            },
            {
              caching: "None",
              diskSizeGB: 200,
              lun: 1,
              managedDisk: { storageAccountType: "Premium_LRS" },
            },
          ],
          diskEncryptionConfiguration: { targets: ["OsDisk", "TemporaryDisk"] },
          imageReference: {
            offer: "WindowsServer",
            publisher: "MicrosoftWindowsServer",
            sku: "2016-Datacenter-SmallDisk",
            version: "latest",
          },
          licenseType: "Windows_Server",
          nodeAgentSkuId: "batch.node.windows amd64",
          nodePlacementConfiguration: { policy: "Zonal" },
          osDisk: { ephemeralOSDiskSettings: { placement: "CacheDisk" } },
          windowsConfiguration: { enableAutomaticUpdates: false },
        },
      },
      networkConfiguration: {
        endpointConfiguration: {
          inboundNatPools: [
            {
              name: "testnat",
              backendPort: 12001,
              frontendPortRangeEnd: 15100,
              frontendPortRangeStart: 15000,
              networkSecurityGroupRules: [
                {
                  access: "Allow",
                  priority: 150,
                  sourceAddressPrefix: "192.100.12.45",
                  sourcePortRanges: ["1", "2"],
                },
                {
                  access: "Deny",
                  priority: 3500,
                  sourceAddressPrefix: "*",
                  sourcePortRanges: ["*"],
                },
              ],
              protocol: "TCP",
            },
          ],
        },
      },
      scaleSettings: {
        autoScale: { evaluationInterval: "PT5M", formula: "$TargetDedicatedNodes=1" },
      },
      vmSize: "STANDARD_D4",
    },
  );
  console.log(result);
}
