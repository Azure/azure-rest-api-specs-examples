from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_network_mappings_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="9112a37f-0f3e-46ec-9c00-060c6edca071",
        resource_group_name="srcBvte2a14C27",
        resource_name="srce2avaultbvtaC27",
    )

    response = client.replication_network_mappings.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationNetworkMappings_List.json
if __name__ == "__main__":
    main()
