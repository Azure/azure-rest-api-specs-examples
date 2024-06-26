from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python virtual_network_function_definition_version_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.network_function_definition_versions.begin_create_or_update(
        resource_group_name="rg",
        publisher_name="TestPublisher",
        network_function_definition_group_name="TestNetworkFunctionDefinitionGroupName",
        network_function_definition_version_name="1.0.0",
        parameters={
            "location": "eastus",
            "properties": {
                "deployParameters": '{"virtualMachineName":{"type":"string"},"extendedLocationName":{"type":"string"},"cpuCores":{"type":"int"},"memorySizeGB":{"type":"int"},"cloudServicesNetworkAttachment":{"type":"object","properties":{"networkAttachmentName":{"type":"string"},"attachedNetworkId":{"type":"string"},"ipAllocationMethod":{"type":"string"},"ipv4Address":{"type":"string"},"ipv6Address":{"type":"string"},"defaultGateway":{"type":"string"}},"required":["attachedNetworkId","ipAllocationMethod"]},"networkAttachments":{"type":"array","items":{"type":"object","properties":{"networkAttachmentName":{"type":"string"},"attachedNetworkId":{"type":"string"},"ipAllocationMethod":{"type":"string"},"ipv4Address":{"type":"string"},"ipv6Address":{"type":"string"},"defaultGateway":{"type":"string"}},"required":["attachedNetworkId","ipAllocationMethod"]}},"storageProfile":{"type":"object","properties":{"osDisk":{"type":"object","properties":{"createOption":{"type":"string"},"deleteOption":{"type":"string"},"diskSizeGB":{"type":"integer"}},"required":["diskSizeGB"]}},"required":["osDisk"]},"sshPublicKeys":{"type":"array","items":{"type":"object","properties":{"keyData":{"type":"string"}},"required":["keyData"]}},"userData":{"type":"string"},"adminUsername":{"type":"string"},"bootMethod":{"type":"string","default":"UEFI","enum":["UEFI","BIOS"]},"isolateEmulatorThread":{"type":"string"},"virtioInterface":{"type":"string"},"placementHints":{"type":"array","items":{"type":"object","properties":{"hintType":{"type":"string","enum":["Affinity","AntiAffinity"]},"resourceId":{"type":"string"},"schedulingExecution":{"type":"string","enum":["Soft","Hard"]},"scope":{"type":"string"}},"required":["hintType","schedulingExecution","resourceId","scope"]}}}',
                "description": "test NFDV for AzureOperatorNexus",
                "networkFunctionTemplate": {
                    "networkFunctionApplications": [
                        {
                            "artifactProfile": {
                                "artifactStore": {
                                    "id": "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"
                                },
                                "imageArtifactProfile": {"imageName": "test-image", "imageVersion": "1.0.0"},
                            },
                            "artifactType": "ImageFile",
                            "dependsOnProfile": {
                                "installDependsOn": [],
                                "uninstallDependsOn": [],
                                "updateDependsOn": [],
                            },
                            "deployParametersMappingRuleProfile": {
                                "applicationEnablement": "Unknown",
                                "imageMappingRuleProfile": {"userConfiguration": ""},
                            },
                            "name": "testImageRole",
                        },
                        {
                            "artifactProfile": {
                                "artifactStore": {
                                    "id": "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore"
                                },
                                "templateArtifactProfile": {
                                    "templateName": "test-template",
                                    "templateVersion": "1.0.0",
                                },
                            },
                            "artifactType": "ArmTemplate",
                            "dependsOnProfile": {
                                "installDependsOn": ["testImageRole"],
                                "uninstallDependsOn": ["testImageRole"],
                                "updateDependsOn": ["testImageRole"],
                            },
                            "deployParametersMappingRuleProfile": {
                                "applicationEnablement": "Unknown",
                                "templateMappingRuleProfile": {
                                    "templateParameters": '{"virtualMachineName":"{deployParameters.virtualMachineName}","extendedLocationName":"{deployParameters.extendedLocationName}","cpuCores":"{deployParameters.cpuCores}","memorySizeGB":"{deployParameters.memorySizeGB}","cloudServicesNetworkAttachment":"{deployParameters.cloudServicesNetworkAttachment}","networkAttachments":"{deployParameters.networkAttachments}","sshPublicKeys":"{deployParameters.sshPublicKeys}","storageProfile":"{deployParameters.storageProfile}","isolateEmulatorThread":"{deployParameters.isolateEmulatorThread}","virtioInterface":"{deployParameters.virtioInterface}","userData":"{deployParameters.userData}","adminUsername":"{deployParameters.adminUsername}","bootMethod":"{deployParameters.bootMethod}","placementHints":"{deployParameters.placementHints}"}'
                                },
                            },
                            "name": "testTemplateRole",
                        },
                    ],
                    "nfviType": "AzureOperatorNexus",
                },
                "networkFunctionType": "VirtualNetworkFunction",
                "versionState": "Preview",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureOperatorNexus/VirtualNetworkFunctionDefinitionVersionCreate.json
if __name__ == "__main__":
    main()
