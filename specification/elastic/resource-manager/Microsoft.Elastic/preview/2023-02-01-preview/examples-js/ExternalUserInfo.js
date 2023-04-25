const { MicrosoftElastic } = require("@azure/arm-elastic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create User inside elastic deployment which are used by customers to perform operations on the elastic deployment
 *
 * @summary Create User inside elastic deployment which are used by customers to perform operations on the elastic deployment
 * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/ExternalUserInfo.json
 */
async function externalUserCreateOrUpdate() {
  const subscriptionId =
    process.env["ELASTIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ELASTIC_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftElastic(credential, subscriptionId);
  const result = await client.externalUser.createOrUpdate(resourceGroupName, monitorName);
  console.log(result);
}
