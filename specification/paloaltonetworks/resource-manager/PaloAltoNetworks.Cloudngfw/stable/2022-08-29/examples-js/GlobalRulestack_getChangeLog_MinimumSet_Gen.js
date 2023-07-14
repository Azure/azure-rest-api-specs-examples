const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get changelog
 *
 * @summary Get changelog
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/GlobalRulestack_getChangeLog_MinimumSet_Gen.json
 */
async function globalRulestackGetChangeLogMinimumSetGen() {
  const globalRulestackName = "praval";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.globalRulestack.getChangeLog(globalRulestackName);
  console.log(result);
}
