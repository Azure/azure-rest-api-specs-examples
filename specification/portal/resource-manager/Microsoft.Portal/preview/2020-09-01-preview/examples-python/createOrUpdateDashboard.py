from azure.identity import DefaultAzureCredential
from azure.mgmt.portal import Portal

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-portal
# USAGE
    python create_or_update_a_dashboard.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = Portal(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.dashboards.create_or_update(
        resource_group_name="testRG",
        dashboard_name="testDashboard",
        dashboard={
            "location": "eastus",
            "properties": {
                "lenses": [
                    {
                        "order": 1,
                        "parts": [
                            {"position": {"colSpan": 3, "rowSpan": 4, "x": 1, "y": 2}},
                            {"position": {"colSpan": 6, "rowSpan": 6, "x": 5, "y": 5}},
                        ],
                    },
                    {"order": 2, "parts": []},
                ],
                "metadata": {"metadata": {"ColSpan": 2, "RowSpan": 1, "X": 4, "Y": 3}},
            },
            "tags": {"aKey": "aValue", "anotherKey": "anotherValue"},
        },
    )
    print(response)


# x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/createOrUpdateDashboard.json
if __name__ == "__main__":
    main()
