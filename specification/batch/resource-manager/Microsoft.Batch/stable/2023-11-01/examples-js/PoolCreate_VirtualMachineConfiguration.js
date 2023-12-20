const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolCreate_VirtualMachineConfiguration.json
 */
async function createPoolFullVirtualMachineConfiguration() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    deploymentConfiguration: {
      virtualMachineConfiguration: {
        dataDisks: [
          {
            caching: "ReadWrite",
            diskSizeGB: 30,
            lun: 0,
            storageAccountType: "Premium_LRS",
          },
          {
            caching: "None",
            diskSizeGB: 200,
            lun: 1,
            storageAccountType: "Standard_LRS",
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
      autoScale: {
        evaluationInterval: "PT5M",
        formula: "$TargetDedicatedNodes=1",
      },
    },
    vmSize: "STANDARD_D4",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.create(
    resourceGroupName,
    accountName,
    poolName,
    parameters,
  );
  console.log(result);
}
