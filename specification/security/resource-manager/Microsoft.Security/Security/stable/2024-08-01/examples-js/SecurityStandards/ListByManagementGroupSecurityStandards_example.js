const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a list of all relevant security standards over a scope
 *
 * @summary get a list of all relevant security standards over a scope
 * x-ms-original-file: 2024-08-01/SecurityStandards/ListByManagementGroupSecurityStandards_example.json
 */
async function listSecurityStandardsByManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (const item of client.securityStandards.list(
    "providers/Microsoft.Management/managementGroups/contoso",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
