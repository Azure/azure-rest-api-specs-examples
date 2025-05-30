from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python vmss_public_ip_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.public_ip_addresses.get_virtual_machine_scale_set_public_ip_address(
        resource_group_name="vmss-tester",
        virtual_machine_scale_set_name="vmss1",
        virtualmachine_index=1,
        network_interface_name="nic1",
        ip_configuration_name="ip1",
        public_ip_address_name="pub1",
    )
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VmssPublicIpGet.json
if __name__ == "__main__":
    main()
