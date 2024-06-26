const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an Azure API Management API if it has been onboarded to Defender for APIs. If an Azure API Management API is onboarded to Defender for APIs, the system will monitor the operations within the Azure API Management API for intrusive behaviors and provide alerts for attacks that have been detected.
 *
 * @summary Gets an Azure API Management API if it has been onboarded to Defender for APIs. If an Azure API Management API is onboarded to Defender for APIs, the system will monitor the operations within the Azure API Management API for intrusive behaviors and provide alerts for attacks that have been detected.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-11-20-preview/examples/ApiCollections/APICollection_Get_example.json
 */
async function getsAnAzureApiManagementApiIfItHasBeenOnboardedToDefenderForApIs() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "3fa85f64-5717-4562-b3fc-2c963f66afa6";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiCollectionId = "echo-api";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.aPICollection.get(resourceGroupName, serviceName, apiCollectionId);
  console.log(result);
}
