from azure.identity import DefaultAzureCredential

from azure.mgmt.avs import AVSClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-avs
# USAGE
    python virtual_machines_restrict_movement.py

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

    client.virtual_machines.begin_restrict_movement(
        resource_group_name="group1",
        private_cloud_name="cloud1",
        cluster_name="cluster1",
        virtual_machine_id="vm-209",
        restrict_movement={"restrictMovement": "Enabled"},
    ).result()


# x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2024-09-01/examples/VirtualMachines_RestrictMovement.json
if __name__ == "__main__":
    main()
