from azure.identity import DefaultAzureCredential
from azure.mgmt.sqlvirtualmachine import SqlVirtualMachineManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sqlvirtualmachine
# USAGE
    python create_or_update_availability_group_listener.py

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

    response = client.availability_group_listeners.begin_create_or_update(
        resource_group_name="testrg",
        sql_virtual_machine_group_name="testvmgroup",
        availability_group_listener_name="agl-test",
        parameters={
            "properties": {
                "availabilityGroupName": "ag-test",
                "loadBalancerConfigurations": [
                    {
                        "loadBalancerResourceId": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb-test",
                        "privateIpAddress": {
                            "ipAddress": "10.1.0.112",
                            "subnetResourceId": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default",
                        },
                        "probePort": 59983,
                        "sqlVirtualMachineInstances": [
                            "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2",
                            "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3",
                        ],
                    }
                ],
                "port": 1433,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/CreateOrUpdateAvailabilityGroupListener.json
if __name__ == "__main__":
    main()
