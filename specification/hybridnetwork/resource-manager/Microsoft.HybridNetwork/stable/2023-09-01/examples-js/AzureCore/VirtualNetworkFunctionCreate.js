const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network function resource.
 *
 * @summary Creates or updates a network function resource.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureCore/VirtualNetworkFunctionCreate.json
 */
async function createVirtualNetworkFunctionResourceOnAzureCore() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg";
  const networkFunctionName = "testNf";
  const parameters = {
    location: "eastus",
    properties: {
      allowSoftwareUpdate: false,
      configurationType: "Open",
      deploymentValues:
        '{"virtualMachineName":"test-VM","cpuCores":4,"memorySizeGB":8,"cloudServicesNetworkAttachment":{"attachedNetworkId":"test-csnet","ipAllocationMethod":"Dynamic","networkAttachmentName":"test-cs-vlan"},"networkAttachments":[{"attachedNetworkId":"test-l3vlan","defaultGateway":"True","ipAllocationMethod":"Dynamic","networkAttachmentName":"test-vlan"}],"sshPublicKeys":[{"keyData":"ssh-rsa CMIIIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0TqlveKKlc2MFvEmuXJiLGBsY1t4ML4uiRADGSZlnc+7Ugv3h+MCjkkwOKiOdsNo8k4KSBIG5GcQfKYOOd17AJvqCL6cGQbaLuqv0a64jeDm8oO8/xN/IM0oKw7rMr/2oAJOgIsfeXPkRxWWic9AVIS++H5Qi2r7bUFX+cqFsyUCAwEBBQ=="}],"storageProfile":{"osDisk":{"createOption":"Ephemeral","deleteOption":"Delete","diskSizeGB":10}},"userData":"testUserData","adminUsername":"testUser","virtioInterface":"Transitional","isolateEmulatorThread":"False","bootMethod":"BIOS","placementHints":[]}',
      networkFunctionDefinitionVersionResourceReference: {
        id: "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1",
        idType: "Open",
      },
      nfviId: "/subscriptions/subid/resourceGroups/testResourceGroup",
      nfviType: "AzureCore",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.networkFunctions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkFunctionName,
    parameters,
  );
  console.log(result);
}
