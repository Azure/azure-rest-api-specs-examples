const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an Internet Gateway rule resource.
 *
 * @summary Creates an Internet Gateway rule resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternetGatewayRules_Create_MaximumSet_Gen.json
 */
async function internetGatewayRulesCreateMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const internetGatewayRuleName = "example-internetGatewayRule";
  const body = {
    annotation: "annotationValue",
    location: "eastus",
    ruleProperties: { action: "Allow", addressList: ["10.10.10.10"] },
    tags: { keyID: "keyValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.internetGatewayRules.beginCreateAndWait(
    resourceGroupName,
    internetGatewayRuleName,
    body
  );
  console.log(result);
}
