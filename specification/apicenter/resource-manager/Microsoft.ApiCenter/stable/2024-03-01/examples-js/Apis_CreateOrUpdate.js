const { AzureAPICenter } = require("@azure/arm-apicenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates new or updates existing API.
 *
 * @summary Creates new or updates existing API.
 * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_CreateOrUpdate.json
 */
async function apisCreateOrUpdate() {
  const subscriptionId =
    process.env["APICENTER_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APICENTER_RESOURCE_GROUP"] || "contoso-resources";
  const serviceName = "contoso";
  const workspaceName = "default";
  const apiName = "echo-api";
  const payload = {
    properties: {
      description: "A simple HTTP request/response service.",
      customProperties: { author: "John Doe" },
      externalDocumentation: [{ title: "Onboarding docs", url: "https://docs.contoso.com" }],
      kind: "rest",
      license: { url: "https://contoso.com/license" },
      lifecycleStage: "design",
      termsOfService: { url: "https://contoso.com/terms-of-service" },
      title: "Echo API",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureAPICenter(credential, subscriptionId);
  const result = await client.apis.createOrUpdate(
    resourceGroupName,
    serviceName,
    workspaceName,
    apiName,
    payload,
  );
  console.log(result);
}
