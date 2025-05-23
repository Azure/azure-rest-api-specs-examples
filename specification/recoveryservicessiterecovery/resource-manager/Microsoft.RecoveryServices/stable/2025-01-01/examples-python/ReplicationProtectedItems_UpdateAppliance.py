from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_protected_items_update_appliance.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="7c943c1b-5122-4097-90c8-861411bdd574",
        resource_group_name="Ayan-0106-SA-RG",
        resource_name="Ayan-0106-SA-Vault",
    )

    response = client.replication_protected_items.begin_update_appliance(
        fabric_name="Ayan-0106-SA-Vaultreplicationfabric",
        protection_container_name="Ayan-0106-SA-Vaultreplicationcontainer",
        replicated_protected_item_name="idclab-vcen67_50158124-8857-3e08-0893-2ddf8ebb8c1f",
        appliance_update_input={
            "properties": {
                "providerSpecificDetails": {"instanceType": "InMageRcm", "runAsAccountId": ""},
                "targetApplianceId": "",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectedItems_UpdateAppliance.json
if __name__ == "__main__":
    main()
