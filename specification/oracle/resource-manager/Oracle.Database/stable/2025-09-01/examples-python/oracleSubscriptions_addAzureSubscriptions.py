from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python oracle_subscriptions_add_azure_subscriptions.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OracleDatabaseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.oracle_subscriptions.begin_add_azure_subscriptions(
        body={"azureSubscriptionIds": ["00000000-0000-0000-0000-000000000001"]},
    ).result()


# x-ms-original-file: 2025-09-01/oracleSubscriptions_addAzureSubscriptions.json
if __name__ == "__main__":
    main()
