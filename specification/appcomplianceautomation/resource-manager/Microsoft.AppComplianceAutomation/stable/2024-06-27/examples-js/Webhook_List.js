const {
  AppComplianceAutomationToolForMicrosoft365,
} = require("@azure/arm-appcomplianceautomation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the AppComplianceAutomation webhook list.
 *
 * @summary Get the AppComplianceAutomation webhook list.
 * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Webhook_List.json
 */
async function webhookList() {
  const skipToken = "1";
  const top = 100;
  const reportName = "testReportName";
  const options = { skipToken, top };
  const credential = new DefaultAzureCredential();
  const client = new AppComplianceAutomationToolForMicrosoft365(credential);
  const resArray = new Array();
  for await (let item of client.webhook.list(reportName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
