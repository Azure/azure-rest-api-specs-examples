const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronize attestation record from app compliance.
 *
 * @summary Synchronize attestation record from app compliance.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_SyncCertRecord.json
 */
async function reportSyncCertRecord() {
  const reportName = "testReportName";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.report.beginSyncCertRecordAndWait(reportName, {
    certRecord: {
      certificationStatus: "CertIngestion",
      controls: [
        {
          controlId: "Operational_Security_10",
          controlStatus: "Approved",
        },
      ],
      ingestionStatus: "EvidenceResubmitted",
      offerGuid: "00000000-0000-0000-0000-000000000001",
    },
  });
  console.log(result);
}
