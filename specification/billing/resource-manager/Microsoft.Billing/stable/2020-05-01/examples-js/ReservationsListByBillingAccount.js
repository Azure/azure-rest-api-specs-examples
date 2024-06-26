const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations for a billing account and the roll up counts of reservations group by provisioning states.
 *
 * @summary Lists the reservations for a billing account and the roll up counts of reservations group by provisioning states.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ReservationsListByBillingAccount.json
 */
async function reservationsListByBillingAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const filter = "properties/reservedResourceType eq 'VirtualMachines'";
  const orderby = "properties/userFriendlyAppliedScopeType asc";
  const selectedState = "Succeeded";
  const options = {
    filter,
    orderby,
    selectedState,
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservations.listByBillingAccount(billingAccountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationsListByBillingAccount().catch(console.error);
