```javascript
const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function checksWhetherNameIsAvailableForCommunicationResource() {
  const subscriptionId = "subid";
  const supportTicketName = "testticket";
  const checkNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Support/communications",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.communications.checkNameAvailability(
    supportTicketName,
    checkNameAvailabilityInput
  );
  console.log(result);
}

checksWhetherNameIsAvailableForCommunicationResource().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-support_2.0.1/sdk/support/arm-support/README.md) on how to add the SDK to your project and authenticate.
