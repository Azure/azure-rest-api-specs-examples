from azure.identity import DefaultAzureCredential
from azure.mgmt.scvmm import SCVMM

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python restart_virtual_machine.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SCVMM(
        credential=DefaultAzureCredential(),
        subscription_id="fd3c3665-1729-4b7b-9a38-238e83b0f98b",
    )

    response = client.virtual_machines.begin_restart(
        resource_group_name="testrg",
        virtual_machine_name="DemoVM",
    ).result()
    print(response)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/RestartVirtualMachine.json
if __name__ == "__main__":
    main()
