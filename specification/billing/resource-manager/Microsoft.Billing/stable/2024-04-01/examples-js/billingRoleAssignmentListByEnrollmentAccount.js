const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the role assignments for the caller on a enrollment account. The operation is supported for billing accounts of type Enterprise Agreement.
 *
 * @summary Lists the role assignments for the caller on a enrollment account. The operation is supported for billing accounts of type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentListByEnrollmentAccount.json
 */
async function billingRoleAssignmentListByEnrollmentAccount() {
  const billingAccountName = "6100092";
  const enrollmentAccountName = "123456";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingRoleAssignments.listByEnrollmentAccount(
    billingAccountName,
    enrollmentAccountName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
