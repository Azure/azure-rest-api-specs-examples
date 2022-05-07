Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-support_2.0.1/sdk/support/arm-support/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function checksWhetherNameIsAvailableForSupportTicketResource() {
  const subscriptionId = "subid";
  const checkNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Support/supportTickets",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.supportTickets.checkNameAvailability(checkNameAvailabilityInput);
  console.log(result);
}

checksWhetherNameIsAvailableForSupportTicketResource().catch(console.error);
```
