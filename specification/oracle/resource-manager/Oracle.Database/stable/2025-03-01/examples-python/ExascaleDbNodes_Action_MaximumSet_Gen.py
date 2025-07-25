from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python exascale_db_nodes_action_maximum_set_gen.py

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

    response = client.exascale_db_nodes.begin_action(
        resource_group_name="rgopenapi",
        exadb_vm_cluster_name="vmClusterName",
        exascale_db_node_name="dbNodeName",
        body={"action": "Start"},
    ).result()
    print(response)


# x-ms-original-file: 2025-03-01/ExascaleDbNodes_Action_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
