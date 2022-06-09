```javascript
const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function listSupportTicketsCreatedOnOrAfterACertainDateAndInOpenStateForASubscription() {
  const subscriptionId = "subid";
  const filter = "createdDate ge 2020-03-10T22:08:51Z and status eq 'Open'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.supportTickets.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSupportTicketsCreatedOnOrAfterACertainDateAndInOpenStateForASubscription().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-support_2.0.1/sdk/support/arm-support/README.md) on how to add the SDK to your project and authenticate.
