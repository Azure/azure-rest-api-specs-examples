const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get student join requests
 *
 * @summary get student join requests
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequestList.json
 */
async function joinRequestList() {
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const includeDenied = false;
  const options = { includeDenied };
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.joinRequests.list(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
