const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates Microsoft Defender for Cloud security operator on the given scope.
 *
 * @summary Creates Microsoft Defender for Cloud security operator on the given scope.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-01-01-preview/examples/SecurityOperators/PutSecurityOperatorByName_example.json
 */
async function createASecurityOperatorOnTheGivenScope() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const pricingName = "CloudPosture";
  const securityOperatorName = "DefenderCSPMSecurityOperator";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityOperators.createOrUpdate(pricingName, securityOperatorName);
  console.log(result);
}
