from azure.identity import DefaultAzureCredential

from azure.mgmt.standbypool import StandbyPoolMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-standbypool
# USAGE
    python standby_virtual_machines_list_by_standby_virtual_machine_pool_resource.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StandbyPoolMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.standby_virtual_machines.list_by_standby_virtual_machine_pool_resource(
        resource_group_name="rgstandbypool",
        standby_virtual_machine_pool_name="pool",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2025-03-01/StandbyVirtualMachines_ListByStandbyVirtualMachinePoolResource.json
if __name__ == "__main__":
    main()
