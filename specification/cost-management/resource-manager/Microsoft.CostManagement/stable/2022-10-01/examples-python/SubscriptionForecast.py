from azure.identity import DefaultAzureCredential
from azure.mgmt.costmanagement import CostManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-costmanagement
# USAGE
    python subscription_forecast.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CostManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.forecast.usage(
        scope="subscriptions/00000000-0000-0000-0000-000000000000",
        parameters={
            "dataset": {
                "aggregation": {"totalCost": {"function": "Sum", "name": "Cost"}},
                "filter": {
                    "and": [
                        {
                            "or": [
                                {
                                    "dimensions": {
                                        "name": "ResourceLocation",
                                        "operator": "In",
                                        "values": ["East US", "West Europe"],
                                    }
                                },
                                {"tags": {"name": "Environment", "operator": "In", "values": ["UAT", "Prod"]}},
                            ]
                        },
                        {"dimensions": {"name": "ResourceGroup", "operator": "In", "values": ["API"]}},
                    ]
                },
                "granularity": "Daily",
            },
            "includeActualCost": False,
            "includeFreshPartialCost": False,
            "timePeriod": {"from": "2022-08-01T00:00:00+00:00", "to": "2022-08-31T23:59:59+00:00"},
            "timeframe": "Custom",
            "type": "Usage",
        },
    )
    print(response)


# x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/SubscriptionForecast.json
if __name__ == "__main__":
    main()
