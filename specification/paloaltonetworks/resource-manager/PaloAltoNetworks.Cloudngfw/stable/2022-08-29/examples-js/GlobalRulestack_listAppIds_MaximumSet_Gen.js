const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of AppIds for GlobalRulestack ApiVersion
 *
 * @summary List of AppIds for GlobalRulestack ApiVersion
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/GlobalRulestack_listAppIds_MaximumSet_Gen.json
 */
async function globalRulestackListAppIdsMaximumSetGen() {
  const globalRulestackName = "praval";
  const appIdVersion = "8543";
  const appPrefix = "pref";
  const skip = "a6a321";
  const top = 20;
  const options = {
    appIdVersion,
    appPrefix,
    skip,
    top,
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.listAppIds(globalRulestackName, options);
  console.log(result);
}
