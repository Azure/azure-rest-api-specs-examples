from azure.identity import DefaultAzureCredential
from azure.mgmt.costmanagement import CostManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-costmanagement
# USAGE
    python external_billing_account_alerts.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CostManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.alerts.list_external(
        external_cloud_provider_type="externalBillingAccounts",
        external_cloud_provider_id="100",
    )
    print(response)


# x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalBillingAccountAlerts.json
if __name__ == "__main__":
    main()
