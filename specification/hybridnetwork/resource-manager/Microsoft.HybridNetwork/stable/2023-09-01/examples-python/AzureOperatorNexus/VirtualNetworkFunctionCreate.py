from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python virtual_network_function_create.py

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

    response = client.network_functions.begin_create_or_update(
        resource_group_name="rg",
        network_function_name="testNf",
        parameters={
            "location": "eastus",
            "properties": {
                "allowSoftwareUpdate": False,
                "configurationType": "Open",
                "deploymentValues": '{"virtualMachineName":"test-VM","extendedLocationName":"test-cluster","cpuCores":4,"memorySizeGB":8,"cloudServicesNetworkAttachment":{"attachedNetworkId":"test-csnet","ipAllocationMethod":"Dynamic","networkAttachmentName":"test-cs-vlan"},"networkAttachments":[{"attachedNetworkId":"test-l3vlan","defaultGateway":"True","ipAllocationMethod":"Dynamic","networkAttachmentName":"test-vlan"}],"sshPublicKeys":[{"keyData":"ssh-rsa CMIIIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0TqlveKKlc2MFvEmuXJiLGBsY1t4ML4uiRADGSZlnc+7Ugv3h+MCjkkwOKiOdsNo8k4KSBIG5GcQfKYOOd17AJvqCL6cGQbaLuqv0a64jeDm8oO8/xN/IM0oKw7rMr/2oAJOgIsfeXPkRxWWic9AVIS++H5Qi2r7bUFX+cqFsyUCAwEBBQ=="}],"storageProfile":{"osDisk":{"createOption":"Ephemeral","deleteOption":"Delete","diskSizeGB":10}},"userData":"testUserData","adminUsername":"testUser","virtioInterface":"Transitional","isolateEmulatorThread":"False","bootMethod":"BIOS","placementHints":[]}',
                "networkFunctionDefinitionVersionResourceReference": {
                    "id": "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1",
                    "idType": "Open",
                },
                "nfviId": "/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation",
                "nfviType": "AzureOperatorNexus",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureOperatorNexus/VirtualNetworkFunctionCreate.json
if __name__ == "__main__":
    main()
