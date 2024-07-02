const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the AppComplianceAutomation report list for the tenant.
 *
 * @summary Get the AppComplianceAutomation report list for the tenant.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_List.json
 */
async function reportList() {
  const skipToken = "1";
  const top = 100;
  const offerGuid = "00000000-0000-0000-0000-000000000000";
  const reportCreatorTenantId = "00000000-0000-0000-0000-000000000000";
  const options = {
    skipToken,
    top,
    offerGuid,
    reportCreatorTenantId,
  };
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const resArray = new Array();
  for await (let item of client.report.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
