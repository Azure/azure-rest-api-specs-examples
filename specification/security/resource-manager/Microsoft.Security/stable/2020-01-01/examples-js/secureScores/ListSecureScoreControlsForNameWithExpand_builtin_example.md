Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all security controls for a specific initiative within a scope
 *
 * @summary Get all security controls for a specific initiative within a scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/ListSecureScoreControlsForNameWithExpand_builtin_example.json
 */
async function getSecurityControlsAndTheirCurrentScoreForTheSpecifiedInitiativeWithTheExpandParameter() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const secureScoreName = "ascScore";
  const expand = "definition";
  const options = {
    expand,
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.secureScoreControls.listBySecureScore(secureScoreName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getSecurityControlsAndTheirCurrentScoreForTheSpecifiedInitiativeWithTheExpandParameter().catch(
  console.error
);
```
