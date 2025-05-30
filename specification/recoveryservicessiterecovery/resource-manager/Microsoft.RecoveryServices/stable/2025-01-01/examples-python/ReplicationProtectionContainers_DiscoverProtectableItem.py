from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_protection_containers_discover_protectable_item.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="7c943c1b-5122-4097-90c8-861411bdd574",
        resource_group_name="MadhaviVRG",
        resource_name="MadhaviVault",
    )

    response = client.replication_protection_containers.begin_discover_protectable_item(
        fabric_name="V2A-W2K12-660",
        protection_container_name="cloud_7328549c-5c37-4459-a3c2-e35f9ef6893c",
        discover_protectable_item_request={
            "properties": {"friendlyName": "Test", "ipAddress": "10.150.2.3", "osType": "Windows"}
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionContainers_DiscoverProtectableItem.json
if __name__ == "__main__":
    main()
