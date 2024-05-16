const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Information protection policies of a specific management group.
 *
 * @summary Information protection policies of a specific management group.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/ListInformationProtectionPolicies_example.json
 */
async function getInformationProtectionPolicies() {
  const scope =
    "providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (let item of client.informationProtectionPolicies.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
