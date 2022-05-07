Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-frontdoor_5.0.1/sdk/frontdoor/arm-frontdoor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new Rules Engine Configuration with the specified name within the specified Front Door.
 *
 * @summary Creates a new Rules Engine Configuration with the specified name within the specified Front Door.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineCreate.json
 */
async function createOrUpdateASpecificRulesEngineConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const frontDoorName = "frontDoor1";
  const rulesEngineName = "rulesEngine1";
  const rulesEngineParameters = {
    rules: [
      {
        name: "Rule1",
        action: {
          routeConfigurationOverride: {
            odataType: "#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration",
            customFragment: "fragment",
            customHost: "www.bing.com",
            customPath: "/api",
            customQueryString: "a=b",
            redirectProtocol: "HttpsOnly",
            redirectType: "Moved",
          },
        },
        matchConditions: [
          {
            rulesEngineMatchValue: ["CH"],
            rulesEngineMatchVariable: "RemoteAddr",
            rulesEngineOperator: "GeoMatch",
          },
        ],
        matchProcessingBehavior: "Stop",
        priority: 1,
      },
      {
        name: "Rule2",
        action: {
          responseHeaderActions: [
            {
              headerActionType: "Overwrite",
              headerName: "Cache-Control",
              value: "public, max-age=31536000",
            },
          ],
        },
        matchConditions: [
          {
            rulesEngineMatchValue: ["jpg"],
            rulesEngineMatchVariable: "RequestFilenameExtension",
            rulesEngineOperator: "Equal",
            transforms: ["Lowercase"],
          },
        ],
        priority: 2,
      },
      {
        name: "Rule3",
        action: {
          routeConfigurationOverride: {
            odataType: "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration",
            backendPool: {
              id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1",
            },
            cacheConfiguration: {
              cacheDuration: "P1DT12H20M30S",
              dynamicCompression: "Disabled",
              queryParameterStripDirective: "StripOnly",
              queryParameters: "a=b,p=q",
            },
            customForwardingPath: undefined,
            forwardingProtocol: "HttpsOnly",
          },
        },
        matchConditions: [
          {
            negateCondition: false,
            rulesEngineMatchValue: ["allowoverride"],
            rulesEngineMatchVariable: "RequestHeader",
            rulesEngineOperator: "Equal",
            selector: "Rules-Engine-Route-Forward",
            transforms: ["Lowercase"],
          },
        ],
        priority: 3,
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.rulesEngines.beginCreateOrUpdateAndWait(
    resourceGroupName,
    frontDoorName,
    rulesEngineName,
    rulesEngineParameters
  );
  console.log(result);
}

createOrUpdateASpecificRulesEngineConfiguration().catch(console.error);
```
