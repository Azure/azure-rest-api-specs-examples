from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python db_nodes_action.py

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

    response = client.db_nodes.begin_action(
        resource_group_name="rg000",
        cloudvmclustername="cluster1",
        dbnodeocid="ocid1....aaaaaa",
        body={"action": "Start"},
    ).result()
    print(response)


# x-ms-original-file: 2025-03-01/dbNodes_action.json
if __name__ == "__main__":
    main()
