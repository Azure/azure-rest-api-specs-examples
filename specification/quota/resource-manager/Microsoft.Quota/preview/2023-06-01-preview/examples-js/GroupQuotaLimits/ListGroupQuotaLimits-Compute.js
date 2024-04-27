const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the GroupQuotaLimits for the all resource for a specific  resourceProvider and $filter passed.
The $filter=location eq {location} is required to location specific resources groupQuota.
 *
 * @summary Gets the GroupQuotaLimits for the all resource for a specific  resourceProvider and $filter passed.
The $filter=location eq {location} is required to location specific resources groupQuota.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimits/ListGroupQuotaLimits-Compute.json
 */
async function groupQuotaLimitsListRequestForCompute() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const resourceProviderName = "Microsoft.Compute";
  const filter = "location eq westus";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.groupQuotaLimits.list(
    managementGroupId,
    groupQuotaName,
    resourceProviderName,
    filter,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
