const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or Update a Http Route Config.
 *
 * @summary Create or Update a Http Route Config.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2025-02-02-preview/examples/HttpRouteConfig_CreateOrUpdate_PathSepPrefix.json
 */
async function createOrUpdateHttpRoutePathSeparatedPrefixRule() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "testcontainerenv";
  const httpRouteName = "httproutefriendlyname";
  const httpRouteConfigEnvelope = {
    properties: {
      customDomains: [
        {
          name: "example.com",
          bindingType: "Disabled",
          certificateId:
            "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-1",
        },
      ],
      rules: [
        {
          description: "random-description",
          routes: [
            {
              action: { prefixRewrite: "/v1/api" },
              match: { caseSensitive: true, pathSeparatedPrefix: "/v1" },
            },
          ],
          targets: [{ containerApp: "capp-1", label: "label-1" }],
        },
      ],
    },
  };
  const options = {
    httpRouteConfigEnvelope,
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.httpRouteConfigOperations.createOrUpdate(
    resourceGroupName,
    environmentName,
    httpRouteName,
    options,
  );
  console.log(result);
}
