const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the AppComplianceAutomation scoping configuration of the specific report.
 *
 * @summary Get the AppComplianceAutomation scoping configuration of the specific report.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ScopingConfiguration_CreateOrUpdate.json
 */
async function scopingConfigurationCreateOrUpdate() {
  const reportName = "testReportName";
  const scopingConfigurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.scopingConfiguration.createOrUpdate(
    reportName,
    scopingConfigurationName,
    {
      properties: {
        answers: [
          {
            answers: ["Azure"],
            questionId: "GEN20_hostingEnvironment",
          },
          {
            answers: [],
            questionId: "DHP_G07_customerDataProcess",
          },
          {
            answers: [],
            questionId: "Tier2InitSub_serviceCommunicate",
          },
        ],
      },
    },
  );
  console.log(result);
}
