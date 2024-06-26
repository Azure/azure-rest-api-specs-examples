const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the instruction by name. These are custom billing instructions and are only applicable for certain customers.
 *
 * @summary Get the instruction by name. These are custom billing instructions and are only applicable for certain customers.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Instruction.json
 */
async function instruction() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const instructionName = "{instructionName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.instructions.get(
    billingAccountName,
    billingProfileName,
    instructionName
  );
  console.log(result);
}

instruction().catch(console.error);
