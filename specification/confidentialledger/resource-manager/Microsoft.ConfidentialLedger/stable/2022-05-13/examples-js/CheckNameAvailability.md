```javascript
const { ConfidentialLedgerClient } = require("@azure/arm-confidentialledger");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to To check whether a resource name is available.
 *
 * @summary To check whether a resource name is available.
 * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/CheckNameAvailability.json
 */
async function checkNameAvailability() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const nameAvailabilityRequest = {
    name: "sample-name",
    type: "Microsoft.ConfidentialLedger/ledgers",
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfidentialLedgerClient(credential, subscriptionId);
  const result = await client.checkNameAvailability(nameAvailabilityRequest);
  console.log(result);
}

checkNameAvailability().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-confidentialledger_1.0.0/sdk/confidentialledger/arm-confidentialledger/README.md) on how to add the SDK to your project and authenticate.
