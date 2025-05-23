from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicesbackup.passivestamp import RecoveryServicesBackupPassiveClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesbackup
# USAGE
    python get_access_token_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesBackupPassiveClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.recovery_points.get_access_token(
        vault_name="rshvault",
        resource_group_name="rshhtestmdvmrg",
        fabric_name="Azure",
        container_name="IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
        protected_item_name="VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall",
        recovery_point_id="26083826328862",
        parameters={
            "properties": {
                "audience": "https://RecoveryServices/IaasCoord/aadmgmt/ase",
                "servicePrincipalObjectId": "5ecd8123-cf74-4037-83e9-9246b227b351",
                "tenantId": "33e01921-4d64-4f8c-a055-5bdaffd5e33d",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-11-15/examples/AzureIaasVm/GetAccessToken_Get.json
if __name__ == "__main__":
    main()
