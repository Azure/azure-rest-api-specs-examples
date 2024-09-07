const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an associated tenant for the billing account.
 *
 * @summary Create or update an associated tenant for the billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/associatedTenantsCreateOrUpdate.json
 */
async function associatedTenantsCreateOrUpdate() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const associatedTenantName = "11111111-1111-1111-1111-111111111111";
  const parameters = {
    properties: {
      billingManagementState: "Active",
      displayName: "Contoso Finance",
      provisioningManagementState: "Pending",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.associatedTenants.beginCreateOrUpdateAndWait(
    billingAccountName,
    associatedTenantName,
    parameters,
  );
  console.log(result);
}
