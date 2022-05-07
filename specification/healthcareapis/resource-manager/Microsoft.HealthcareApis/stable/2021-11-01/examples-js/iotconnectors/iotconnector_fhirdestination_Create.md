Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-healthcareapis_2.1.0/sdk/healthcareapis/arm-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an IoT Connector FHIR destination resource with the specified parameters.
 *
 * @summary Creates or updates an IoT Connector FHIR destination resource with the specified parameters.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Create.json
 */
async function createOrUpdateAnIotConnectorFhirDestination() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const iotConnectorName = "blue";
  const fhirDestinationName = "dest1";
  const iotFhirDestination = {
    fhirMapping: {
      content: {
        template: [
          {
            template: {
              codes: [
                {
                  code: "8867-4",
                  display: "Heart rate",
                  system: "http://loinc.org",
                },
              ],
              periodInterval: 60,
              typeName: "heartrate",
              value: {
                defaultPeriod: 5000,
                unit: "count/min",
                valueName: "hr",
                valueType: "SampledData",
              },
            },
            templateType: "CodeValueFhir",
          },
        ],
        templateType: "CollectionFhirTemplate",
      },
    },
    fhirServiceResourceId:
      "subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice",
    location: "westus",
    resourceIdentityResolutionType: "Create",
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.iotConnectorFhirDestination.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    iotConnectorName,
    fhirDestinationName,
    iotFhirDestination
  );
  console.log(result);
}

createOrUpdateAnIotConnectorFhirDestination().catch(console.error);
```
