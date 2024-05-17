const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Details of a specific Compliance.
 *
 * @summary Details of a specific Compliance.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliance_example.json
 */
async function getSecurityComplianceDataForADay() {
  const scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const complianceName = "2018-01-01Z";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.compliances.get(scope, complianceName);
  console.log(result);
}
