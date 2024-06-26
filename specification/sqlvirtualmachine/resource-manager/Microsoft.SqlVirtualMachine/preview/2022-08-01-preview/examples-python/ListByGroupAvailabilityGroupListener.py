from azure.identity import DefaultAzureCredential
from azure.mgmt.sqlvirtualmachine import SqlVirtualMachineManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sqlvirtualmachine
# USAGE
    python list_by_group_availability_group_listener.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlVirtualMachineManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.availability_group_listeners.list_by_group(
        resource_group_name="testrg",
        sql_virtual_machine_group_name="testvmgroup",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/ListByGroupAvailabilityGroupListener.json
if __name__ == "__main__":
    main()
