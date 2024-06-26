from azure.identity import DefaultAzureCredential
from azure.mgmt.maps import AzureMapsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-maps
# USAGE
    python get_maps_creator.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureMapsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="21a9967a-e8a9-4656-a70b-96ff1c4d05a0",
    )

    response = client.creators.get(
        resource_group_name="myResourceGroup",
        account_name="myMapsAccount",
        creator_name="myCreator",
    )
    print(response)


# x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/GetMapsCreator.json
if __name__ == "__main__":
    main()
