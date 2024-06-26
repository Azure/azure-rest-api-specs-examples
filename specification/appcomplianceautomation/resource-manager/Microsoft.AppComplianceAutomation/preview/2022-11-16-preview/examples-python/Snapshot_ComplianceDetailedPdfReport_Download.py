from azure.identity import DefaultAzureCredential
from azure.mgmt.appcomplianceautomation import AppComplianceAutomationToolForMicrosoft365

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcomplianceautomation
# USAGE
    python snapshot_download_compliance_detailed_pdf_report.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AppComplianceAutomationToolForMicrosoft365(
        credential=DefaultAzureCredential(),
    )

    response = client.snapshot.begin_download(
        report_name="testReportName",
        snapshot_name="testSnapshotName",
        parameters={
            "downloadType": "ComplianceDetailedPdfReport",
            "offerGuid": "00000000-0000-0000-0000-000000000000",
            "reportCreatorTenantId": "00000000-0000-0000-0000-000000000000",
        },
    ).result()
    print(response)


# x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_ComplianceDetailedPdfReport_Download.json
if __name__ == "__main__":
    main()
