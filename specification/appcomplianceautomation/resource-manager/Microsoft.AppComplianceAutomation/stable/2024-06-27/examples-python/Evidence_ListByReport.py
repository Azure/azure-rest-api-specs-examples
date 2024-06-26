from azure.identity import DefaultAzureCredential

from azure.mgmt.appcomplianceautomation import AppComplianceAutomationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcomplianceautomation
# USAGE
    python evidence_list_by_report.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AppComplianceAutomationMgmtClient(
        credential=DefaultAzureCredential(),
    )

    response = client.evidence.list_by_report(
        report_name="reportName",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Evidence_ListByReport.json
if __name__ == "__main__":
    main()
