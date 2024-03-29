const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List secure scores for all your Microsoft Defender for Cloud initiatives within your current scope.
 *
 * @summary List secure scores for all your Microsoft Defender for Cloud initiatives within your current scope.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/ListSecureScores_example.json
 */
async function listSecureScores() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.secureScores.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
