from azure.identity import DefaultAzureCredential
from azure.mgmt.datashare import DataShareManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datashare
# USAGE
    python accounts_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataShareManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="433a8dfd-e5d5-4e77-ad86-90acdc75eb1a",
    )

    response = client.accounts.begin_create(
        resource_group_name="SampleResourceGroup",
        account_name="Account1",
        account={
            "identity": {"type": "SystemAssigned"},
            "location": "West US 2",
            "tags": {"tag1": "Red", "tag2": "White"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/Accounts_Create.json
if __name__ == "__main__":
    main()
