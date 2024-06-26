from azure.identity import DefaultAzureCredential
from azure.mgmt.datalake.store import DataLakeStoreAccountManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datalake-store
# USAGE
    python locations_get_usage.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataLakeStoreAccountManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.locations.get_usage(
        location="WestUS",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Locations_GetUsage.json
if __name__ == "__main__":
    main()
