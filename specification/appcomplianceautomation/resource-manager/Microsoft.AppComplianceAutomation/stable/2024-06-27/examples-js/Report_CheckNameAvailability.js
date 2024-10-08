const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if the given name is available for a report.
 *
 * @summary Check if the given name is available for a report.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_CheckNameAvailability.json
 */
async function reportCheckNameAvailability() {
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.providerActions.checkNameAvailability({
    name: "reportABC",
    type: "Microsoft.AppComplianceAutomation/reports",
  });
  console.log(result);
}
