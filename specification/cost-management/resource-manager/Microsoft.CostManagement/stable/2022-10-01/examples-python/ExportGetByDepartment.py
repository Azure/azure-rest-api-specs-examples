from azure.identity import DefaultAzureCredential
from azure.mgmt.costmanagement import CostManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-costmanagement
# USAGE
    python export_get_by_department.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CostManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.exports.get(
        scope="providers/Microsoft.Billing/billingAccounts/12/departments/1234",
        export_name="TestExport",
    )
    print(response)


# x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportGetByDepartment.json
if __name__ == "__main__":
    main()
