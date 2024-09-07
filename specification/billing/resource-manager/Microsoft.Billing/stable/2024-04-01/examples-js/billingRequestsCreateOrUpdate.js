const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a billing request.
 *
 * @summary Create or update a billing request.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRequestsCreateOrUpdate.json
 */
async function billingRequestsCreateOrUpdate() {
  const billingRequestName = "00000000-0000-0000-0000-000000000000";
  const parameters = {
    properties: {
      type: "RoleAssignment",
      additionalInformation: { roleId: "40000000-aaaa-bbbb-cccc-200000000006" },
      decisionReason: "New team member",
      requestScope:
        "/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx",
      status: "Pending",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingRequests.beginCreateOrUpdateAndWait(
    billingRequestName,
    parameters,
  );
  console.log(result);
}
