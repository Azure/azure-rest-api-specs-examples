const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the availability of a resource name. This API should be used to check the uniqueness of the name for adding a new communication to the support ticket.
 *
 * @summary Check the availability of a resource name. This API should be used to check the uniqueness of the name for adding a new communication to the support ticket.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CheckNameAvailabilityForNoSubscriptionSupportTicketCommunication.json
 */
async function checksWhetherNameIsAvailableForCommunicationResource() {
  const subscriptionId =
    process.env["SUPPORT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const supportTicketName = "testticket";
  const checkNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Support/communications",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.communicationsNoSubscription.checkNameAvailability(
    supportTicketName,
    checkNameAvailabilityInput
  );
  console.log(result);
}
