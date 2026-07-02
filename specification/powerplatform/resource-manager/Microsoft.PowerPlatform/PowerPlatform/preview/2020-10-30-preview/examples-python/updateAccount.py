from azure.identity import DefaultAzureCredential

from azure.mgmt.powerplatform import PowerPlatformMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-powerplatform
# USAGE
    python update_account.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PowerPlatformMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.accounts.update(
        account_name="account",
        resource_group_name="resourceGroup",
        parameters={
            "location": "East US",
            "properties": {"description": "Description of account."},
            "tags": {"Organization": "Administration"},
        },
    )
    print(response)


# x-ms-original-file: 2020-10-30-preview/updateAccount.json
if __name__ == "__main__":
    main()
