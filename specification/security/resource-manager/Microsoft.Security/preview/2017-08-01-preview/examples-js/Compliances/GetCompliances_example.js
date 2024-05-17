const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Compliance scores of the specific management group.
 *
 * @summary The Compliance scores of the specific management group.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliances_example.json
 */
async function getSecurityComplianceDataOverTime() {
  const scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (let item of client.compliances.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
