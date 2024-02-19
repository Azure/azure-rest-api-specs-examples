const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Summaries resource.
 *
 * @summary Gets the Summaries resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/Summaries_Get_MaximumSet_Gen.json
 */
async function summariesGetMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "libzegdqkcxmhqhhhcxm";
  const resourceGroupName =
    process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootdiscovery";
  const siteName = "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps";
  const summaryName = "vjB";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.summaries.get(resourceGroupName, siteName, summaryName);
  console.log(result);
}
