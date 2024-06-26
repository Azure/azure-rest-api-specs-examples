const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get secure score for a specific Microsoft Defender for Cloud initiative within your current scope. For the ASC Default initiative, use 'ascScore'.
 *
 * @summary Get secure score for a specific Microsoft Defender for Cloud initiative within your current scope. For the ASC Default initiative, use 'ascScore'.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/GetSecureScoresSingle_example.json
 */
async function getSingleSecureScore() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const secureScoreName = "ascScore";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.secureScores.get(secureScoreName);
  console.log(result);
}
