```javascript
const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a FHIR Service resource with the specified parameters.
 *
 * @summary Creates or updates a FHIR Service resource with the specified parameters.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_Create.json
 */
async function createOrUpdateAFhirService() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const fhirServiceName = "fhirservice1";
  const fhirservice = {
    accessPolicies: [
      { objectId: "c487e7d1-3210-41a3-8ccc-e9372b78da47" },
      { objectId: "5b307da8-43d4-492b-8b66-b0294ade872f" },
    ],
    acrConfiguration: { loginServers: ["test1.azurecr.io"] },
    authenticationConfiguration: {
      audience: "https://azurehealthcareapis.com",
      authority: "https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc",
      smartProxyEnabled: true,
    },
    corsConfiguration: {
      allowCredentials: false,
      headers: ["*"],
      maxAge: 1440,
      methods: ["DELETE", "GET", "OPTIONS", "PATCH", "POST", "PUT"],
      origins: ["*"],
    },
    exportConfiguration: { storageAccountName: "existingStorageAccount" },
    identity: { type: "SystemAssigned" },
    kind: "fhir-R4",
    location: "westus",
    tags: {
      additionalProp1: "string",
      additionalProp2: "string",
      additionalProp3: "string",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.fhirServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    fhirServiceName,
    fhirservice
  );
  console.log(result);
}

createOrUpdateAFhirService().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-healthcareapis_2.1.0/sdk/healthcareapis/arm-healthcareapis/README.md) on how to add the SDK to your project and authenticate.
