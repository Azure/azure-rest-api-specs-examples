from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_protected_items_update_mobility_service.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="b364ed8d-4279-4bf8-8fd1-56f8fa0ae05c",
        resource_group_name="wcusValidations",
        resource_name="WCUSVault",
    )

    response = client.replication_protected_items.begin_update_mobility_service(
        fabric_name="WIN-JKKJ31QI8U2",
        protection_container_name="cloud_c6780228-83bd-4f3e-a70e-cb46b7da33a0",
        replicated_protected_item_name="79dd20ab-2b40-11e7-9791-0050568f387e",
        update_mobility_service_request={"properties": {"runAsAccountId": "2"}},
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectedItems_UpdateMobilityService.json
if __name__ == "__main__":
    main()
