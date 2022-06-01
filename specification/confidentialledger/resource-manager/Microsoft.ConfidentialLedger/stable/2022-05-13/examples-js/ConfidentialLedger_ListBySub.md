Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-confidentialledger_1.0.0/sdk/confidentialledger/arm-confidentialledger/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConfidentialLedgerClient } = require("@azure/arm-confidentialledger");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the properties of all Confidential Ledgers.
 *
 * @summary Retrieves the properties of all Confidential Ledgers.
 * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/ConfidentialLedger_ListBySub.json
 */
async function confidentialLedgerListBySub() {
  const subscriptionId = "0000000-0000-0000-0000-000000000001";
  const credential = new DefaultAzureCredential();
  const client = new ConfidentialLedgerClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.ledger.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

confidentialLedgerListBySub().catch(console.error);
```
