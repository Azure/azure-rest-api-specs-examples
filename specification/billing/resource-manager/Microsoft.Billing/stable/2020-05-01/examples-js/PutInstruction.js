const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an instruction. These are custom billing instructions and are only applicable for certain customers.
 *
 * @summary Creates or updates an instruction. These are custom billing instructions and are only applicable for certain customers.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInstruction.json
 */
async function putInstruction() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const instructionName = "{instructionName}";
  const parameters = {
    amount: 5000,
    endDate: new Date("2020-12-30T21:26:47.997Z"),
    startDate: new Date("2019-12-30T21:26:47.997Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.instructions.put(
    billingAccountName,
    billingProfileName,
    instructionName,
    parameters
  );
  console.log(result);
}

putInstruction().catch(console.error);
