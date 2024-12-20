from azure.identity import DefaultAzureCredential

from azure.mgmt.extendedlocation import CustomLocations

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-extendedlocation
# USAGE
    python custom_locations_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CustomLocations(
        credential=DefaultAzureCredential(),
        subscription_id="11111111-2222-3333-4444-555555555555",
    )

    client.custom_locations.begin_delete(
        resource_group_name="testresourcegroup",
        resource_name="customLocation01",
    ).result()


# x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsDelete.json
if __name__ == "__main__":
    main()
