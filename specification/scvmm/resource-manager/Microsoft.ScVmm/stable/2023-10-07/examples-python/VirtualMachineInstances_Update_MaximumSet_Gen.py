from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.scvmm import ScVmmMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python virtual_machine_instances_update_maximum_set_gen.py

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

    response = client.virtual_machine_instances.begin_update(
        resource_uri="gtgclehcbsyave",
        properties={
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
                    "limitCpuForMigration": "true",
                    "memoryMB": 5,
                },
                "infrastructureProfile": {"checkpointType": "jkbpzjxpeegackhsvikrnlnwqz"},
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
                "storageProfile": {
                    "disks": [
                        {
                            "bus": 8,
                            "busType": "zu",
                            "diskId": "ltdrwcfjklpsimhzqyh",
                            "diskSizeGB": 30,
                            "lun": 10,
                            "name": "fgnckfymwdsqnfxkdvexuaobe",
                            "storageQoSPolicy": {"id": "o", "name": "ceiyfrflu"},
                            "vhdType": "cnbeeeylrvopigdynvgpkfp",
                        }
                    ]
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
