from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python protection_containers_inquire_result.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.protection_container_operation_results.get(
        vault_name="testvault",
        resource_group_name="test-rg",
        fabric_name="Azure",
        container_name="VMAppContainer;Compute;testRG;testSQL",
        operation_id="00000000-0000-0000-0000-000000000000",
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureStorage/ProtectionContainers_Inquire_Result.json
if __name__ == "__main__":
    main()
