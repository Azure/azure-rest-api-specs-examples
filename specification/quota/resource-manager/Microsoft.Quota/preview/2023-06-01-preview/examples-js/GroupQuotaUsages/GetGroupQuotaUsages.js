const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the GroupQuotas usages and limits(quota). Location is required paramter.
 *
 * @summary Gets the GroupQuotas usages and limits(quota). Location is required paramter.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaUsages/GetGroupQuotaUsages.json
 */
async function groupQuotasUsagesList() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const resourceProviderName = "Microsoft.Compute";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.groupQuotaUsages.list(
    managementGroupId,
    groupQuotaName,
    resourceProviderName,
    location,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
