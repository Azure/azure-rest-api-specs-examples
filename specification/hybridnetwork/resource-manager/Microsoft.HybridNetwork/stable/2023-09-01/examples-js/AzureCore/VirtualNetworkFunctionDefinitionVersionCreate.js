const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network function definition version.
 *
 * @summary Creates or updates a network function definition version.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureCore/VirtualNetworkFunctionDefinitionVersionCreate.json
 */
async function createOrUpdateANetworkFunctionDefinitionVersionResourceForAzureCoreVnf() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg";
  const publisherName = "TestPublisher";
  const networkFunctionDefinitionGroupName = "TestNetworkFunctionDefinitionGroupName";
  const networkFunctionDefinitionVersionName = "1.0.0";
  const parameters = {
    location: "eastus",
    properties: {
      description: "test NFDV for AzureCore",
      deployParameters:
        '{"virtualMachineName":{"type":"string"},"cpuCores":{"type":"int"},"memorySizeGB":{"type":"int"},"cloudServicesNetworkAttachment":{"type":"object","properties":{"networkAttachmentName":{"type":"string"},"attachedNetworkId":{"type":"string"},"ipAllocationMethod":{"type":"string"},"ipv4Address":{"type":"string"},"ipv6Address":{"type":"string"},"defaultGateway":{"type":"string"}},"required":["attachedNetworkId","ipAllocationMethod"]},"networkAttachments":{"type":"array","items":{"type":"object","properties":{"networkAttachmentName":{"type":"string"},"attachedNetworkId":{"type":"string"},"ipAllocationMethod":{"type":"string"},"ipv4Address":{"type":"string"},"ipv6Address":{"type":"string"},"defaultGateway":{"type":"string"}},"required":["attachedNetworkId","ipAllocationMethod"]}},"storageProfile":{"type":"object","properties":{"osDisk":{"type":"object","properties":{"createOption":{"type":"string"},"deleteOption":{"type":"string"},"diskSizeGB":{"type":"integer"}},"required":["diskSizeGB"]}},"required":["osDisk"]},"sshPublicKeys":{"type":"array","items":{"type":"object","properties":{"keyData":{"type":"string"}},"required":["keyData"]}},"userData":{"type":"string"},"adminUsername":{"type":"string"},"bootMethod":{"type":"string","default":"UEFI","enum":["UEFI","BIOS"]},"isolateEmulatorThread":{"type":"string"},"virtioInterface":{"type":"string"},"placementHints":{"type":"array","items":{"type":"object","properties":{"hintType":{"type":"string","enum":["Affinity","AntiAffinity"]},"resourceId":{"type":"string"},"schedulingExecution":{"type":"string","enum":["Soft","Hard"]},"scope":{"type":"string"}},"required":["hintType","schedulingExecution","resourceId","scope"]}}}',
      networkFunctionTemplate: {
        networkFunctionApplications: [
          {
            name: "testImageRole",
            artifactProfile: {
              artifactStore: {
                id: "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore",
              },
              vhdArtifactProfile: { vhdName: "test-image", vhdVersion: "1-0-0" },
            },
            artifactType: "VhdImageFile",
            dependsOnProfile: {
              installDependsOn: [],
              uninstallDependsOn: [],
              updateDependsOn: [],
            },
            deployParametersMappingRuleProfile: {
              applicationEnablement: "Unknown",
              vhdImageMappingRuleProfile: { userConfiguration: "" },
            },
          },
          {
            name: "testTemplateRole",
            artifactProfile: {
              artifactStore: {
                id: "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore",
              },
              templateArtifactProfile: {
                templateName: "test-template",
                templateVersion: "1.0.0",
              },
            },
            artifactType: "ArmTemplate",
            dependsOnProfile: {
              installDependsOn: ["testImageRole"],
              uninstallDependsOn: ["testImageRole"],
              updateDependsOn: ["testImageRole"],
            },
            deployParametersMappingRuleProfile: {
              applicationEnablement: "Unknown",
              templateMappingRuleProfile: {
                templateParameters:
                  '{"virtualMachineName":"{deployParameters.virtualMachineName}","cpuCores":"{deployParameters.cpuCores}","memorySizeGB":"{deployParameters.memorySizeGB}","cloudServicesNetworkAttachment":"{deployParameters.cloudServicesNetworkAttachment}","networkAttachments":"{deployParameters.networkAttachments}","sshPublicKeys":"{deployParameters.sshPublicKeys}","storageProfile":"{deployParameters.storageProfile}","isolateEmulatorThread":"{deployParameters.isolateEmulatorThread}","virtioInterface":"{deployParameters.virtioInterface}","userData":"{deployParameters.userData}","adminUsername":"{deployParameters.adminUsername}","bootMethod":"{deployParameters.bootMethod}","placementHints":"{deployParameters.placementHints}"}',
              },
            },
          },
        ],
        nfviType: "AzureCore",
      },
      networkFunctionType: "VirtualNetworkFunction",
      versionState: "Preview",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.networkFunctionDefinitionVersions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    publisherName,
    networkFunctionDefinitionGroupName,
    networkFunctionDefinitionVersionName,
    parameters,
  );
  console.log(result);
}
