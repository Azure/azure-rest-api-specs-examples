from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python recovery_points_get.py

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

    response = client.recovery_points.get(
        vault_name="rshvault",
        resource_group_name="rshhtestmdvmrg",
        fabric_name="Azure",
        container_name="IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
        protected_item_name="VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
        recovery_point_id="26083826328862",
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/RecoveryPoints_Get.json
if __name__ == "__main__":
    main()
