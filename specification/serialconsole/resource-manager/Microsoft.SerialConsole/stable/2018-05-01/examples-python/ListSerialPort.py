from azure.identity import DefaultAzureCredential
from azure.mgmt.serialconsole import MicrosoftSerialConsoleClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-serialconsole
# USAGE
    python list_serial_ports_for_parent_resources.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftSerialConsoleClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-00000-0000-0000-000000000000",
    )

    response = client.serial_ports.list(
        resource_group_name="myResourceGroup",
        resource_provider_namespace="Microsoft.Compute",
        parent_resource_type="virtualMachines",
        parent_resource="myVM",
    )
    print(response)


# x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/ListSerialPort.json
if __name__ == "__main__":
    main()
