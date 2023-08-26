const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates new or updates existing specified API of the API Management service instance.
 *
 * @summary Creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateSoapPassThroughApiUsingWsdlImport.json
 */
async function apiManagementCreateSoapPassThroughApiUsingWsdlImport() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "soapApi";
  const parameters = {
    format: "wsdl-link",
    path: "currency",
    soapApiType: "soap",
    value: "http://www.webservicex.net/CurrencyConvertor.asmx?WSDL",
    wsdlSelector: {
      wsdlEndpointName: "CurrencyConvertorSoap",
      wsdlServiceName: "CurrencyConvertor",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    apiId,
    parameters
  );
  console.log(result);
}
