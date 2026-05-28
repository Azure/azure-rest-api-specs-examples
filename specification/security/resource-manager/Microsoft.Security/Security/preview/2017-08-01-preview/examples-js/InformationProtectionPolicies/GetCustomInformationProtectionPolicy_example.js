const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to details of the information protection policy.
 *
 * @summary details of the information protection policy.
 * x-ms-original-file: 2017-08-01-preview/InformationProtectionPolicies/GetCustomInformationProtectionPolicy_example.json
 */
async function getTheCustomizedInformationProtectionPolicyForAManagementGroup() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.informationProtectionPolicies.get(
    "providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e",
    "custom",
  );
  console.log(result);
}
