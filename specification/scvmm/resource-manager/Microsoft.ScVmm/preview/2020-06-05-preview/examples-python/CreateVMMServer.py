from azure.identity import DefaultAzureCredential
from azure.mgmt.scvmm import SCVMM

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python create_vmm_server.py

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

    response = client.vmm_servers.begin_create_or_update(
        resource_group_name="testrg",
        vmm_server_name="ContosoVMMServer",
        body={
            "extendedLocation": {
                "name": "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso",
                "type": "customLocation",
            },
            "location": "East US",
            "properties": {
                "credentials": {"password": "password", "username": "testuser"},
                "fqdn": "VMM.contoso.com",
                "port": 1234,
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateVMMServer.json
if __name__ == "__main__":
    main()
