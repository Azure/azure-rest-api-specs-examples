from azure.identity import DefaultAzureCredential
from azure.mgmt.alertsmanagement import AlertsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-alertsmanagement
# USAGE
    python alerts_change_state.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AlertsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="9e261de7-c804-4b9d-9ebf-6f50fe350a9a",
    )

    response = client.alerts.change_state(
        alert_id="66114d64-d9d9-478b-95c9-b789d6502100",
        new_state="Acknowledged",
    )
    print(response)


# x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_ChangeState.json
if __name__ == "__main__":
    main()
