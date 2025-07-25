from azure.identity import DefaultAzureCredential

from azure.mgmt.mongodbatlas import MongoDBAtlasMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-mongodbatlas
# USAGE
    python organizations_list_by_subscription_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MongoDBAtlasMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.organizations.list_by_subscription()
    for item in response:
        print(item)


# x-ms-original-file: 2025-06-01/Organizations_ListBySubscription_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
