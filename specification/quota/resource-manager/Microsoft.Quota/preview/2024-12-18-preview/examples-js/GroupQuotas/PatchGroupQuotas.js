const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the GroupQuotas for the name passed. A GroupQuotas RequestId will be returned by the Service. The status can be polled periodically. The status Async polling is using standards defined at - https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/async-api-reference.md#asynchronous-operations. Use the OperationsStatus URI provided in Azure-AsyncOperation header, the duration will be specified in retry-after header. Once the operation gets to terminal state - Succeeded | Failed, then the URI will change to Get URI and full details can be checked.
 Any change in the filters will be applicable to the future quota assignments, existing quota allocated to subscriptions from the GroupQuotas remains unchanged.
 *
 * @summary Updates the GroupQuotas for the name passed. A GroupQuotas RequestId will be returned by the Service. The status can be polled periodically. The status Async polling is using standards defined at - https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/async-api-reference.md#asynchronous-operations. Use the OperationsStatus URI provided in Azure-AsyncOperation header, the duration will be specified in retry-after header. Once the operation gets to terminal state - Succeeded | Failed, then the URI will change to Get URI and full details can be checked.
 Any change in the filters will be applicable to the future quota assignments, existing quota allocated to subscriptions from the GroupQuotas remains unchanged.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/GroupQuotas/PatchGroupQuotas.json
 */
async function groupQuotasPatchRequestForCompute() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const groupQuotasPatchRequestBody = {
    properties: { displayName: "UpdatedGroupQuota1" },
  };
  const options = {
    groupQuotasPatchRequestBody,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const result = await client.groupQuotas.beginUpdateAndWait(
    managementGroupId,
    groupQuotaName,
    options,
  );
  console.log(result);
}
