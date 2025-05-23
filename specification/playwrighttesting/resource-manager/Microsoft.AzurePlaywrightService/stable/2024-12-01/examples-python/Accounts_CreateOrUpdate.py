from azure.identity import DefaultAzureCredential

from azure.mgmt.playwrighttesting import PlaywrightTestingMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-playwrighttesting
# USAGE
    python accounts_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PlaywrightTestingMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.accounts.begin_create_or_update(
        resource_group_name="dummyrg",
        account_name="myPlaywrightAccount",
        resource={"location": "westus", "properties": {"regionalAffinity": "Enabled"}, "tags": {"Team": "Dev Exp"}},
    ).result()
    print(response)


# x-ms-original-file: 2024-12-01/Accounts_CreateOrUpdate.json
if __name__ == "__main__":
    main()
