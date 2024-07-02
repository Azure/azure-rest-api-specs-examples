const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the count of reports.
 *
 * @summary Get the count of reports.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_GetCollectionCount.json
 */
async function reportGetCollectionCount() {
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.providerActions.getCollectionCount({
    type: "Microsoft.AppComplianceAutomation/reports",
  });
  console.log(result);
}
