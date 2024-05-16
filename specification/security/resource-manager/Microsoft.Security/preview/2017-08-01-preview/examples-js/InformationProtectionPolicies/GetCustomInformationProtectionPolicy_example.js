const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Details of the information protection policy.
 *
 * @summary Details of the information protection policy.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/GetCustomInformationProtectionPolicy_example.json
 */
async function getTheCustomizedInformationProtectionPolicyForAManagementGroup() {
  const scope =
    "providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e";
  const informationProtectionPolicyName = "custom";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.informationProtectionPolicies.get(
    scope,
    informationProtectionPolicyName,
  );
  console.log(result);
}
