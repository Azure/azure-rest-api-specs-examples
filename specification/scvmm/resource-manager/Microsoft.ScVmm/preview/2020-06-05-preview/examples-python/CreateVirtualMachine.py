from azure.identity import DefaultAzureCredential
from azure.mgmt.scvmm import SCVMM

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python create_virtual_machine.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SCVMM(
        credential=DefaultAzureCredential(),
        subscription_id="fd3c3665-1729-4b7b-9a38-238e83b0f98b",
    )

    response = client.virtual_machines.begin_create_or_update(
        resource_group_name="testrg",
        virtual_machine_name="DemoVM",
        body={
            "extendedLocation": {
                "name": "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso",
                "type": "customLocation",
            },
            "location": "East US",
            "properties": {
                "cloudId": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/Clouds/HRCloud",
                "hardwareProfile": {"cpuCount": 4, "memoryMB": 4096},
                "templateId": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VirtualMachineTemplates/HRVirtualMachineTemplate",
                "vmmServerId": "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateVirtualMachine.json
if __name__ == "__main__":
    main()
