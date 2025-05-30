from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python target_compute_sizes_list_by_replication_protected_items.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="6808dbbc-98c7-431f-a1b1-9580902423b7",
        resource_group_name="avraiMgDiskVaultRG",
        resource_name="avraiMgDiskVault",
    )

    response = client.target_compute_sizes.list_by_replication_protected_items(
        fabric_name="asr-a2a-default-centraluseuap",
        protection_container_name="asr-a2a-default-centraluseuap-container",
        replicated_protected_item_name="468c912d-b1ab-4ea2-97eb-4b5095155db2",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/TargetComputeSizes_ListByReplicationProtectedItems.json
if __name__ == "__main__":
    main()
