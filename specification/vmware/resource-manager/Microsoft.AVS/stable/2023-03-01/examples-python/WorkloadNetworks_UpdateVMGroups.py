from azure.identity import DefaultAzureCredential
from azure.mgmt.avs import AVSClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-avs
# USAGE
    python workload_networks_update_vm_groups.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AVSClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.workload_networks.begin_update_vm_group(
        resource_group_name="group1",
        private_cloud_name="cloud1",
        vm_group_id="vmGroup1",
        workload_network_vm_group={"properties": {"members": ["564d43da-fefc-2a3b-1d92-42855622fa50"], "revision": 1}},
    ).result()
    print(response)


# x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_UpdateVMGroups.json
if __name__ == "__main__":
    main()
