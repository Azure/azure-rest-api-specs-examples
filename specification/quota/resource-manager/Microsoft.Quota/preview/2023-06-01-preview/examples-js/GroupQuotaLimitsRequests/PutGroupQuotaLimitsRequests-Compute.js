const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Put the GroupQuota requests for a specific ResourceProvider/Location/Resource. the location and resourceName ("name": {"value" : "resourceName") properties are specified in the request body. Only 1 resource quota can be requested.
Use the polling API - OperationsStatus URI specified in Azure-AsyncOperation header field, with retry-after duration in seconds to check the intermediate status. This API provides the finals status with the request details and status.
 *
 * @summary Put the GroupQuota requests for a specific ResourceProvider/Location/Resource. the location and resourceName ("name": {"value" : "resourceName") properties are specified in the request body. Only 1 resource quota can be requested.
Use the polling API - OperationsStatus URI specified in Azure-AsyncOperation header field, with retry-after duration in seconds to check the intermediate status. This API provides the finals status with the request details and status.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimitsRequests/PutGroupQuotaLimitsRequests-Compute.json
 */
async function groupQuotaLimitsRequestsCreateOrUpdate() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const resourceProviderName = "Microsoft.Compute";
  const resourceName = "standardav2family";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const result = await client.groupQuotaLimitsRequest.beginCreateOrUpdateAndWait(
    managementGroupId,
    groupQuotaName,
    resourceProviderName,
    resourceName,
  );
  console.log(result);
}
