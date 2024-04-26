const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns only the list of the Azure regions settings, where the GroupQuotas enforcement is enabled. The locations not included in GroupQuota Enforcement will not be listed, the regions in failed status with listed as status Failed.
 *
 * @summary Returns only the list of the Azure regions settings, where the GroupQuotas enforcement is enabled. The locations not included in GroupQuota Enforcement will not be listed, the regions in failed status with listed as status Failed.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasEnforcement/ListGroupQuotaEnforcement.json
 */
async function groupQuotaEnforcementList() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const resourceProviderName = "Microsoft.Compute";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.groupQuotaLocationSettings.list(
    managementGroupId,
    groupQuotaName,
    resourceProviderName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
