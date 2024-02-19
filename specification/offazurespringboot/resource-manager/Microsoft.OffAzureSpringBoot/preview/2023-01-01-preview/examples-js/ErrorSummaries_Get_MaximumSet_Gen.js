const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the ErrorSummaries resource.
 *
 * @summary Gets the ErrorSummaries resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/ErrorSummaries_Get_MaximumSet_Gen.json
 */
async function errorSummariesGetMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "libzegdqkcxmhqhhhcxm";
  const resourceGroupName =
    process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootdiscovery";
  const siteName = "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps";
  const errorSummaryName = "K2lv";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.errorSummaries.get(resourceGroupName, siteName, errorSummaryName);
  console.log(result);
}
