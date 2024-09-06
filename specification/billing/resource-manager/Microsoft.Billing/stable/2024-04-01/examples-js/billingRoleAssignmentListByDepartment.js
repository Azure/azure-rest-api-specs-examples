const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the role assignments for the caller on a department. The operation is supported for billing accounts of type Enterprise Agreement.
 *
 * @summary Lists the role assignments for the caller on a department. The operation is supported for billing accounts of type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentListByDepartment.json
 */
async function billingRoleAssignmentListByDepartment() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30";
  const departmentName = "12345";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.billingRoleAssignments.listByDepartment(
    billingAccountName,
    departmentName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
