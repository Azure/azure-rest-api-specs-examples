from azure.identity import DefaultAzureCredential
from azure.mgmt.connectedvmware import ConnectedVMwareMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-connectedvmware
# USAGE
    python delete_vm_instance_guest_agent.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConnectedVMwareMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.vm_instance_guest_agents.begin_delete(
        resource_uri="subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM",
    ).result()


# x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteVMInstanceGuestAgent.json
if __name__ == "__main__":
    main()
