const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the ErrorSummaries resource in springbootsites.
 *
 * @summary Lists the ErrorSummaries resource in springbootsites.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/ErrorSummaries_ListBySite_MaximumSet_Gen.json
 */
async function errorSummariesListBySiteMaximumSetGen() {
  const subscriptionId =
    process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "libzegdqkcxmhqhhhcxm";
  const resourceGroupName =
    process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootdiscovery";
  const siteName = "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps";
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.errorSummaries.listBySite(resourceGroupName, siteName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
