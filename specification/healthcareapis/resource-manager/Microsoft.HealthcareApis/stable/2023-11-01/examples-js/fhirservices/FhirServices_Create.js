const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a FHIR Service resource with the specified parameters.
 *
 * @summary Creates or updates a FHIR Service resource with the specified parameters.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/fhirservices/FhirServices_Create.json
 */
async function createOrUpdateAFhirService() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HEALTHCAREAPIS_RESOURCE_GROUP"] || "testRG";
  const workspaceName = "workspace1";
  const fhirServiceName = "fhirservice1";
  const fhirservice = {
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
    encryption: {
      customerManagedKeyEncryption: {
        keyEncryptionKeyUrl: "https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion",
      },
    },
    exportConfiguration: { storageAccountName: "existingStorageAccount" },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/subid/resourcegroups/testRG/providers/MicrosoftManagedIdentity/userAssignedIdentities/testMi":
          {},
      },
    },
    implementationGuidesConfiguration: { usCoreMissingData: false },
    importConfiguration: {
      enabled: false,
      initialImportMode: false,
      integrationDataStore: "existingStorageAccount",
    },
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
    fhirservice,
  );
  console.log(result);
}
