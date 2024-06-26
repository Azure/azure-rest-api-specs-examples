from azure.identity import DefaultAzureCredential
from azure.mgmt.alertsmanagement import AlertsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-alertsmanagement
# USAGE
    python get_prometheus_rule_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AlertsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7",
    )

    response = client.prometheus_rule_groups.get(
        resource_group_name="giladstest",
        rule_group_name="myPrometheusRuleGroup",
    )
    print(response)


# x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2021-07-22-preview/examples/getPrometheusRuleGroup.json
if __name__ == "__main__":
    main()
