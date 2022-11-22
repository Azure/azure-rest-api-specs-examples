from azure.identity import DefaultAzureCredential
from azure.mgmt.datamigration import DataMigrationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datamigration
# USAGE
    python usages_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataMigrationManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="90fb80a6-0f71-4761-8f03-921e7396f3c0",
    )

    response = client.usages.list(
        location="westus",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/Usages_List.json
if __name__ == "__main__":
    main()
