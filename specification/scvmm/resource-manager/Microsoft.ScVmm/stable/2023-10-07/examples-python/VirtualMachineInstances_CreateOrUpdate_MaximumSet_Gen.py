from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.scvmm import ScVmmMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python virtual_machine_instances_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ScVmmMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.virtual_machine_instances.begin_create_or_update(
        resource_uri="gtgclehcbsyave",
        resource={
            "extendedLocation": {
                "name": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName",
                "type": "customLocation",
            },
            "properties": {
                "availabilitySets": [
                    {
                        "id": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/availabilitySets/availabilitySetResourceName",
                        "name": "lwbhaseo",
                    }
                ],
                "hardwareProfile": {
                    "cpuCount": 22,
                    "dynamicMemoryEnabled": "true",
                    "dynamicMemoryMaxMB": 2,
                    "dynamicMemoryMinMB": 30,
                    "isHighlyAvailable": "true",
                    "limitCpuForMigration": "true",
                    "memoryMB": 5,
                },
                "infrastructureProfile": {
                    "biosGuid": "xixivxifyql",
                    "checkpointType": "jkbpzjxpeegackhsvikrnlnwqz",
                    "cloudId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/clouds/cloudResourceName",
                    "generation": 28,
                    "inventoryItemId": "ihkkqmg",
                    "lastRestoredVMCheckpoint": {
                        "checkpointID": "wsqmrje",
                        "description": "qurzfrgyflrh",
                        "name": "keqn",
                        "parentCheckpointID": "hqhhzikoxunuqguouw",
                    },
                    "templateId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName",
                    "uuid": "hrpw",
                    "vmName": "qovpayfydhcvfrhe",
                    "vmmServerId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName",
                },
                "networkProfile": {
                    "networkInterfaces": [
                        {
                            "ipv4AddressType": "Dynamic",
                            "ipv6AddressType": "Dynamic",
                            "macAddress": "oaeqqegt",
                            "macAddressType": "Dynamic",
                            "name": "kvofzqulbjlbtt",
                            "nicId": "roxpsvlo",
                            "virtualNetworkId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName",
                        }
                    ]
                },
                "osProfile": {
                    "adminPassword": "vavtppmmhlspydtkzxda",
                    "computerName": "uuxpcxuxcufllc",
                    "osType": "Windows",
                },
                "storageProfile": {
                    "disks": [
                        {
                            "bus": 8,
                            "busType": "zu",
                            "createDiffDisk": "true",
                            "diskId": "ltdrwcfjklpsimhzqyh",
                            "diskSizeGB": 30,
                            "lun": 10,
                            "name": "fgnckfymwdsqnfxkdvexuaobe",
                            "storageQoSPolicy": {"id": "o", "name": "ceiyfrflu"},
                            "templateDiskId": "lcdwrokpyvekqccclf",
                            "vhdType": "cnbeeeylrvopigdynvgpkfp",
                        }
                    ]
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
