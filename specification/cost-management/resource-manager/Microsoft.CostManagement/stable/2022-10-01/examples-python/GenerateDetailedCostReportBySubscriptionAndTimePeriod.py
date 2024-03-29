from azure.identity import DefaultAzureCredential
from azure.mgmt.costmanagement import CostManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-costmanagement
# USAGE
    python generate_detailed_cost_report_by_subscription_and_time_period.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CostManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.generate_detailed_cost_report.begin_create_operation(
        scope="subscriptions/00000000-0000-0000-0000-000000000000",
        parameters={"metric": "ActualCost", "timePeriod": {"end": "2020-03-15", "start": "2020-03-01"}},
    ).result()
    print(response)


# x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/GenerateDetailedCostReportBySubscriptionAndTimePeriod.json
if __name__ == "__main__":
    main()
