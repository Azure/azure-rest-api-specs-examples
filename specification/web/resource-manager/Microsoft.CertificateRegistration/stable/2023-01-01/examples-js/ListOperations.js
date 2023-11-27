const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Implements Csm operations Api to exposes the list of available Csm Apis under the resource provider
 *
 * @summary Description for Implements Csm operations Api to exposes the list of available Csm Apis under the resource provider
 * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-01-01/examples/ListOperations.json
 */
async function listOperations() {
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.certificateRegistrationProvider.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}
