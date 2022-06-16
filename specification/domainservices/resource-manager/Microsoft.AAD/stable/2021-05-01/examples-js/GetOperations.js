const { DomainServicesResourceProvider } = require("@azure/arm-domainservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available Domain Services operations.
 *
 * @summary Lists all the available Domain Services operations.
 * x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/GetOperations.json
 */
async function getOperations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DomainServicesResourceProvider(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.domainServiceOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getOperations().catch(console.error);
