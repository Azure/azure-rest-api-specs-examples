const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the quota allocation request status for the subscriptionId by allocationId.
 *
 * @summary Get the quota allocation request status for the subscriptionId by allocationId.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/SubscriptionQuotaAllocationRequests/SubscriptionQuotaAllocationRequests_Get-Compute.json
 */
async function subscriptionQuotaAllocationRequestsGetRequestForCompute() {
  const subscriptionId =
    process.env["QUOTA_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const resourceProviderName = "Microsoft.Compute";
  const allocationId = "AE000000-0000-0000-0000-00000000000A";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential, subscriptionId);
  const result = await client.groupQuotaSubscriptionAllocationRequest.get(
    managementGroupId,
    groupQuotaName,
    resourceProviderName,
    allocationId,
  );
  console.log(result);
}
