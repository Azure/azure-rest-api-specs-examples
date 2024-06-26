from azure.identity import DefaultAzureCredential
from azure.mgmt.costmanagement import CostManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-costmanagement
# USAGE
    python export_create_or_update_by_management_group.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CostManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.exports.create_or_update(
        scope="providers/Microsoft.Management/managementGroups/TestMG",
        export_name="TestExport",
        parameters={
            "properties": {
                "definition": {
                    "dataSet": {
                        "configuration": {"columns": ["Date", "MeterId", "ResourceId", "ResourceLocation", "Quantity"]},
                        "granularity": "Daily",
                    },
                    "timeframe": "MonthToDate",
                    "type": "ActualCost",
                },
                "deliveryInfo": {
                    "destination": {
                        "container": "exports",
                        "resourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Storage/storageAccounts/ccmeastusdiag182",
                        "rootFolderPath": "ad-hoc",
                    }
                },
                "format": "Csv",
                "schedule": {
                    "recurrence": "Weekly",
                    "recurrencePeriod": {"from": "2020-06-01T00:00:00Z", "to": "2020-10-31T00:00:00Z"},
                    "status": "Active",
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportCreateOrUpdateByManagementGroup.json
if __name__ == "__main__":
    main()
