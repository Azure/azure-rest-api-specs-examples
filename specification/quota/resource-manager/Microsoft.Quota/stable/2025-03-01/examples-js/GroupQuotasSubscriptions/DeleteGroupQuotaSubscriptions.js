const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Removes the subscription from GroupQuotas. The request's TenantId is validated against the subscription's TenantId.
 *
 * @summary Removes the subscription from GroupQuotas. The request's TenantId is validated against the subscription's TenantId.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/GroupQuotasSubscriptions/DeleteGroupQuotaSubscriptions.json
 */
async function groupQuotaSubscriptionsDeleteSubscriptions() {
  const subscriptionId =
    process.env["QUOTA_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential, subscriptionId);
  const result = await client.groupQuotaSubscriptions.beginDeleteAndWait(
    managementGroupId,
    groupQuotaName,
  );
  console.log(result);
}
