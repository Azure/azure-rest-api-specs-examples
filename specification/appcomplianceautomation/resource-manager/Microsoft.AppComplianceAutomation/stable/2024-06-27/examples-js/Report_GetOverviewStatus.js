const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the resource overview status.
 *
 * @summary Get the resource overview status.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_GetOverviewStatus.json
 */
async function reportGetOverviewStatus() {
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const result = await client.providerActions.getOverviewStatus({
    type: "Microsoft.AppComplianceAutomation/reports",
  });
  console.log(result);
}
