const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolCreate_CloudServiceConfiguration.json
 */
async function createPoolFullCloudServiceConfiguration() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    applicationLicenses: ["app-license0", "app-license1"],
    applicationPackages: [
      {
        id: "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234",
        version: "asdf",
      },
    ],
    certificates: [
      {
        id: "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567",
        storeLocation: "LocalMachine",
        storeName: "MY",
        visibility: ["RemoteUser"],
      },
    ],
    deploymentConfiguration: {
      cloudServiceConfiguration: {
        osFamily: "4",
        osVersion: "WA-GUEST-OS-4.45_201708-01",
      },
    },
    displayName: "my-pool-name",
    interNodeCommunication: "Enabled",
    metadata: [
      { name: "metadata-1", value: "value-1" },
      { name: "metadata-2", value: "value-2" },
    ],
    networkConfiguration: {
      publicIPAddressConfiguration: {
        ipAddressIds: [
          "/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135",
          "/subscriptions/subid2/resourceGroups/rg24/providers/Microsoft.Network/publicIPAddresses/ip268",
        ],
        provision: "UserManaged",
      },
      subnetId:
        "/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123",
    },
    scaleSettings: {
      fixedScale: {
        nodeDeallocationOption: "TaskCompletion",
        resizeTimeout: "PT8M",
        targetDedicatedNodes: 6,
        targetLowPriorityNodes: 28,
      },
    },
    startTask: {
      commandLine: "cmd /c SET",
      environmentSettings: [{ name: "MYSET", value: "1234" }],
      maxTaskRetryCount: 6,
      resourceFiles: [
        {
          fileMode: "777",
          filePath: "c:\\temp\\gohere",
          httpUrl: "https://testaccount.blob.core.windows.net/example-blob-file",
        },
      ],
      userIdentity: { autoUser: { elevationLevel: "Admin", scope: "Pool" } },
      waitForSuccess: true,
    },
    taskSchedulingPolicy: { nodeFillType: "Pack" },
    taskSlotsPerNode: 13,
    userAccounts: [
      {
        name: "username1",
        elevationLevel: "Admin",
        linuxUserConfiguration: {
          gid: 4567,
          sshPrivateKey: "sshprivatekeyvalue",
          uid: 1234,
        },
        password: "<ExamplePassword>",
      },
    ],
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
