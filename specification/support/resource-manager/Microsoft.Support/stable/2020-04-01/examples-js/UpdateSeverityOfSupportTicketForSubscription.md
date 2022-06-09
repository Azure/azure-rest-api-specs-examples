```javascript
const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This API allows you to update the severity level, ticket status, and your contact information in the support ticket.<br/><br/>Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer. In such a case, contact your support engineer to request severity update by adding a new communication using the Communications API.<br/><br/>Changing the ticket status to _closed_ is allowed only on an unassigned case. When an engineer is actively working on the ticket, send your ticket closure request by sending a note to your engineer.
 *
 * @summary This API allows you to update the severity level, ticket status, and your contact information in the support ticket.<br/><br/>Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer. In such a case, contact your support engineer to request severity update by adding a new communication using the Communications API.<br/><br/>Changing the ticket status to _closed_ is allowed only on an unassigned case. When an engineer is actively working on the ticket, send your ticket closure request by sending a note to your engineer.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/UpdateSeverityOfSupportTicketForSubscription.json
 */
async function updateSeverityOfASupportTicket() {
  const subscriptionId = "subid";
  const supportTicketName = "testticket";
  const updateSupportTicket = { severity: "critical" };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.supportTickets.update(supportTicketName, updateSupportTicket);
  console.log(result);
}

updateSeverityOfASupportTicket().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-support_2.0.1/sdk/support/arm-support/README.md) on how to add the SDK to your project and authenticate.
