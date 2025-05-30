from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python delete_resource_guard_proxy.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupClient(
        credential=DefaultAzureCredential(),
        subscription_id="0b352192-dcac-4cc7-992e-a96190ccc68c",
    )

    client.resource_guard_proxy.delete(
        vault_name="sampleVault",
        resource_group_name="SampleResourceGroup",
        resource_guard_proxy_name="swaggerExample",
    )


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/ResourceGuardProxyCRUD/DeleteResourceGuardProxy.json
if __name__ == "__main__":
    main()
