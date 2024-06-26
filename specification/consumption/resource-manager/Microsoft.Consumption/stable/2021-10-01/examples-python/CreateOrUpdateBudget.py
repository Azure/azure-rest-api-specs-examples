from azure.identity import DefaultAzureCredential
from azure.mgmt.consumption import ConsumptionManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-consumption
# USAGE
    python create_or_update_budget.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConsumptionManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.budgets.create_or_update(
        scope="subscriptions/00000000-0000-0000-0000-000000000000",
        budget_name="TestBudget",
        parameters={
            "eTag": '"1d34d016a593709"',
            "properties": {
                "amount": 100.65,
                "category": "Cost",
                "filter": {
                    "and": [
                        {
                            "dimensions": {
                                "name": "ResourceId",
                                "operator": "In",
                                "values": [
                                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2",
                                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1",
                                ],
                            }
                        },
                        {"tags": {"name": "category", "operator": "In", "values": ["Dev", "Prod"]}},
                        {"tags": {"name": "department", "operator": "In", "values": ["engineering", "sales"]}},
                    ]
                },
                "notifications": {
                    "Actual_GreaterThan_80_Percent": {
                        "contactEmails": ["johndoe@contoso.com", "janesmith@contoso.com"],
                        "contactGroups": [
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"
                        ],
                        "contactRoles": ["Contributor", "Reader"],
                        "enabled": True,
                        "locale": "en-us",
                        "operator": "GreaterThan",
                        "threshold": 80,
                        "thresholdType": "Actual",
                    }
                },
                "timeGrain": "Monthly",
                "timePeriod": {"endDate": "2018-10-31T00:00:00Z", "startDate": "2017-10-01T00:00:00Z"},
            },
        },
    )
    print(response)


# x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreateOrUpdateBudget.json
if __name__ == "__main__":
    main()
