from azure.identity import DefaultAzureCredential
from azure.mgmt.alertsmanagement import AlertsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-alertsmanagement
# USAGE
    python smart_groups_history.py

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

    response = client.smart_groups.get_history(
        smart_group_id="a808445e-bb38-4751-85c2-1b109ccc1059",
    )
    print(response)


# x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_History.json
if __name__ == "__main__":
    main()
