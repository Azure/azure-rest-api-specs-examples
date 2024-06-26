const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations for a billing profile and the roll up counts of reservations group by provisioning state.
 *
 * @summary Lists the reservations for a billing profile and the roll up counts of reservations group by provisioning state.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ReservationsListByBillingProfile.json
 */
async function reservationsListByBillingProfile() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
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
  for await (let item of client.reservations.listByBillingProfile(
    billingAccountName,
    billingProfileName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationsListByBillingProfile().catch(console.error);
