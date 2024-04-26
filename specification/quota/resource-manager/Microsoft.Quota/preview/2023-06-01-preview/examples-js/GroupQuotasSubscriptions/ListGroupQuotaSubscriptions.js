const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of the subscriptionIds associated with the GroupQuotas.
 *
 * @summary Returns a list of the subscriptionIds associated with the GroupQuotas.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasSubscriptions/ListGroupQuotaSubscriptions.json
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
