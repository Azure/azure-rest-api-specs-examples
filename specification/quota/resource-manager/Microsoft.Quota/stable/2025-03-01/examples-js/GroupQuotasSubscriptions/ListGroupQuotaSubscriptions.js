const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns a list of the subscriptionIds associated with the GroupQuotas.
 *
 * @summary Returns a list of the subscriptionIds associated with the GroupQuotas.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/GroupQuotasSubscriptions/ListGroupQuotaSubscriptions.json
 */
async function groupQuotaSubscriptionsListSubscriptions() {
  const managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
  const groupQuotaName = "groupquota1";
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const resArray = new Array();
  for await (let item of client.groupQuotaSubscriptions.list(managementGroupId, groupQuotaName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
