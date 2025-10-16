from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python vm_clusters_add_vms.py

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

    response = client.cloud_vm_clusters.begin_add_vms(
        resource_group_name="rg000",
        cloudvmclustername="cluster1",
        body={"dbServers": ["ocid1..aaaa", "ocid1..aaaaaa"]},
    ).result()
    print(response)


# x-ms-original-file: 2025-09-01/vmClusters_addVms.json
if __name__ == "__main__":
    main()
