from azure.identity import DefaultAzureCredential
from azure.mgmt.serialconsole import MicrosoftSerialConsoleClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-serialconsole
# USAGE
    python get_the_serial_console_disabled_status_for_a_subscription.py

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

    response = client.get_console_status(
        default="default",
    )
    print(response)


# x-ms-original-file: specification/serialconsole/resource-manager/Microsoft.SerialConsole/stable/2018-05-01/examples/SerialConsoleStatus.json
if __name__ == "__main__":
    main()
