const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enables the GroupQuotas enforcement for the resource provider and the location specified. The resource provider will start using the group quotas as the overall quota for the subscriptions included in the GroupQuota. The subscriptions cannot request quota at subscription level.
The subscriptions share the GroupQuotaLimits assigned to the GroupQuota. If the GroupQuotaLimits is used, then submit a groupQuotaLimit request for the specific resource - provider/location/resource.
Once the GroupQuota Enforcement is enabled then, it cannot be deleted or reverted back. To disable GroupQuota Enforcement -
 1. Remove all the subscriptions from the groupQuota using the delete API for Subscriptions (Check the example - GroupQuotaSubscriptions_Delete).
 2. Ten delete the GroupQuota (Check the example - GroupQuotas_Delete).
 *
 * @summary Enables the GroupQuotas enforcement for the resource provider and the location specified. The resource provider will start using the group quotas as the overall quota for the subscriptions included in the GroupQuota. The subscriptions cannot request quota at subscription level.
The subscriptions share the GroupQuotaLimits assigned to the GroupQuota. If the GroupQuotaLimits is used, then submit a groupQuotaLimit request for the specific resource - provider/location/resource.
Once the GroupQuota Enforcement is enabled then, it cannot be deleted or reverted back. To disable GroupQuota Enforcement -
 1. Remove all the subscriptions from the groupQuota using the delete API for Subscriptions (Check the example - GroupQuotaSubscriptions_Delete).
 2. Ten delete the GroupQuota (Check the example - GroupQuotas_Delete).
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasEnforcement/PutGroupQuotaEnforcement.json
 */
async function groupQuotaLocationSettingsCreateOrUpdate() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const resourceProviderName = "Microsoft.Compute";
  const location = "eastus";
  const locationSettings = {
    properties: { enforcementEnabled: "Enabled" },
  };
  const options = {
    locationSettings,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const result = await client.groupQuotaLocationSettings.beginCreateOrUpdateAndWait(
    managementGroupId,
    groupQuotaName,
    resourceProviderName,
    location,
    options,
  );
  console.log(result);
}
