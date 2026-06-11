const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates new or updates existing specified API of the API Management service instance.
 *
 * @summary creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateSoapPassThroughApiUsingWsdlImport.json
 */
async function apiManagementCreateSoapPassThroughApiUsingWsdlImport() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.createOrUpdate("rg1", "apimService1", "soapApi", {
    format: "wsdl-link",
    path: "currency",
    soapApiType: "soap",
    value: "http://www.webservicex.net/CurrencyConvertor.asmx?WSDL",
    wsdlSelector: {
      wsdlEndpointName: "CurrencyConvertorSoap",
      wsdlServiceName: "CurrencyConvertor",
    },
  });
  console.log(result);
}
