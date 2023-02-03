const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get student join requests
 *
 * @summary get student join requests
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequest.json
 */
async function joinRequest() {
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const joinRequestName = "{joinRequestName}";
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const result = await client.joinRequests.get(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    joinRequestName
  );
  console.log(result);
}
