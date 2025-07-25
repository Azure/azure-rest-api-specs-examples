from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python autonomous_database_generate_wallet.py

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

    response = client.autonomous_databases.generate_wallet(
        resource_group_name="rg000",
        autonomousdatabasename="databasedb1",
        body={"generateType": "Single", "isRegional": False, "password": "********"},
    )
    print(response)


# x-ms-original-file: 2025-03-01/autonomousDatabase_generateWallet.json
if __name__ == "__main__":
    main()
