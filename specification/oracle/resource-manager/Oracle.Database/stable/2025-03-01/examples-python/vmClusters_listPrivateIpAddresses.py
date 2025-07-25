from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python vm_clusters_list_private_ip_addresses.py

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

    response = client.cloud_vm_clusters.list_private_ip_addresses(
        resource_group_name="rg000",
        cloudvmclustername="cluster1",
        body={"subnetId": "ocid1..aaaaaa", "vnicId": "ocid1..aaaaa"},
    )
    print(response)


# x-ms-original-file: 2025-03-01/vmClusters_listPrivateIpAddresses.json
if __name__ == "__main__":
    main()
