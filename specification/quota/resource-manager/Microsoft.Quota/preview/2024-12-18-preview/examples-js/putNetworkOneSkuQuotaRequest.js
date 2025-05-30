const { AzureQuotaExtensionAPI } = require("@azure/arm-quota");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update the quota limit for the specified resource with the requested value. To update the quota, follow these steps:
1. Use the GET operation for quotas and usages to determine how much quota remains for the specific resource and to calculate the new quota limit. These steps are detailed in [this example](https://techcommunity.microsoft.com/t5/azure-governance-and-management/using-the-new-quota-rest-api/ba-p/2183670).
2. Use this PUT operation to update the quota limit. Please check the URI in location header for the detailed status of the request.
 *
 * @summary Create or update the quota limit for the specified resource with the requested value. To update the quota, follow these steps:
1. Use the GET operation for quotas and usages to determine how much quota remains for the specific resource and to calculate the new quota limit. These steps are detailed in [this example](https://techcommunity.microsoft.com/t5/azure-governance-and-management/using-the-new-quota-rest-api/ba-p/2183670).
2. Use this PUT operation to update the quota limit. Please check the URI in location header for the detailed status of the request.
 * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/putNetworkOneSkuQuotaRequest.json
 */
async function quotasPutRequestForNetwork() {
  const resourceName = "MinPublicIpInterNetworkPrefixLength";
  const scope =
    "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus";
  const createQuotaRequest = {
    properties: {
      name: { value: "MinPublicIpInterNetworkPrefixLength" },
      limit: { limitObjectType: "LimitValue", value: 10 },
      resourceType: "MinPublicIpInterNetworkPrefixLength",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureQuotaExtensionAPI(credential);
  const result = await client.quota.beginCreateOrUpdateAndWait(
    resourceName,
    scope,
    createQuotaRequest,
  );
  console.log(result);
}
