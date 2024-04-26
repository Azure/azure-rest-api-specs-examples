const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get API to check the status of a GroupQuota request by requestId.
 *
 * @summary Get API to check the status of a GroupQuota request by requestId.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimitsRequests/GroupQuotaLimitsRequests_Get.json
 */
async function groupQuotaLimitsRequestsGet() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const requestId = "requestId";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const result = await client.groupQuotaLimitsRequest.get(
    managementGroupId,
    groupQuotaName,
    requestId,
  );
  console.log(result);
}
