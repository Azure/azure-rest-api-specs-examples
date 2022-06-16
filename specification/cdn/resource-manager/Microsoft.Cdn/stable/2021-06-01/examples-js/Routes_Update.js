const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing route with the specified route name under the specified subscription, resource group, profile, and AzureFrontDoor endpoint.
 *
 * @summary Updates an existing route with the specified route name under the specified subscription, resource group, profile, and AzureFrontDoor endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Routes_Update.json
 */
async function routesUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const routeName = "route1";
  const routeUpdateProperties = {
    cacheConfiguration: {
      compressionSettings: {
        contentTypesToCompress: ["text/html", "application/octet-stream"],
        isCompressionEnabled: true,
      },
      queryStringCachingBehavior: "IgnoreQueryString",
    },
    customDomains: [
      {
        id: "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/customDomains/domain1",
      },
    ],
    enabledState: "Enabled",
    forwardingProtocol: "MatchRequest",
    httpsRedirect: "Enabled",
    linkToDefaultDomain: "Enabled",
    originGroup: {
      id: "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/originGroups/originGroup1",
    },
    originPath: undefined,
    patternsToMatch: ["/*"],
    ruleSets: [
      {
        id: "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/ruleSets/ruleSet1",
      },
    ],
    supportedProtocols: ["Https", "Http"],
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.routes.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    routeName,
    routeUpdateProperties
  );
  console.log(result);
}

routesUpdate().catch(console.error);
