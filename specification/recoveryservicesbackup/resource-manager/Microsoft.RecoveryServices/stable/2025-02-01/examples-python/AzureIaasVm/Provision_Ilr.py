from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.activestamp import RecoveryServicesBackupClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python provision_ilr.py

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

    client.item_level_recovery_connections.provision(
        vault_name="PySDKBackupTestRsVault",
        resource_group_name="PythonSDKBackupTestRg",
        fabric_name="Azure",
        container_name="iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
        protected_item_name="vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1",
        recovery_point_id="1",
        parameters={
            "properties": {
                "initiatorName": "Hello World",
                "objectType": "IaasVMILRRegistrationRequest",
                "recoveryPointId": "38823086363464",
                "renewExistingRegistration": True,
                "virtualMachineId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/pysdktestrg/providers/Microsoft.Compute/virtualMachines/pysdktestv2vm1",
            }
        },
    )


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/Provision_Ilr.json
if __name__ == "__main__":
    main()
