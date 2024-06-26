const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Azure API Management APIs that have been onboarded to Microsoft Defender for APIs. If an Azure API Management API is onboarded to Microsoft Defender for APIs, the system will monitor the operations within the Azure API Management API for intrusive behaviors and provide alerts for attacks that have been detected.
 *
 * @summary Gets a list of Azure API Management APIs that have been onboarded to Microsoft Defender for APIs. If an Azure API Management API is onboarded to Microsoft Defender for APIs, the system will monitor the operations within the Azure API Management API for intrusive behaviors and provide alerts for attacks that have been detected.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/APICollections_ListByAzureApiManagementService_example.json
 */
async function getsAListOfAzureApiManagementApIsThatHaveBeenOnboardedToMicrosoftDefenderForApis() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "3fa85f64-5717-4562-b3fc-2c963f66afa6";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.aPICollections.listByAzureApiManagementService(
    resourceGroupName,
    serviceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
