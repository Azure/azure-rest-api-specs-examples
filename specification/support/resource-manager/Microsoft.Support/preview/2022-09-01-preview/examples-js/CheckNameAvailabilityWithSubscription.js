const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the availability of a resource name. This API should be used to check the uniqueness of the name for support ticket creation for the selected subscription.
 *
 * @summary Check the availability of a resource name. This API should be used to check the uniqueness of the name for support ticket creation for the selected subscription.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CheckNameAvailabilityWithSubscription.json
 */
async function checksWhetherNameIsAvailableForSupportTicketResource() {
  const subscriptionId = process.env["SUPPORT_SUBSCRIPTION_ID"] || "subid";
  const checkNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Support/supportTickets",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.supportTickets.checkNameAvailability(checkNameAvailabilityInput);
  console.log(result);
}
