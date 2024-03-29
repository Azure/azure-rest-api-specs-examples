from azure.identity import DefaultAzureCredential
from azure.mgmt.datadog import MicrosoftDatadogClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datadog
# USAGE
    python refresh_set_password_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftDatadogClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.monitors.refresh_set_password_link(
        resource_group_name="myResourceGroup",
        monitor_name="myMonitor",
    )
    print(response)


# x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/RefreshSetPassword_Get.json
if __name__ == "__main__":
    main()
