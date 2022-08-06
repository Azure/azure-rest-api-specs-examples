const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Automanage AAD first party Application Service Principal details for the subscription id.
 *
 * @summary Get the Automanage AAD first party Application Service Principal details for the subscription id.
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listServicePrincipalBySubscription.json
 */
async function listServicePrincipalBySubscription() {
  const subscriptionId = "mySubscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.servicePrincipals.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listServicePrincipalBySubscription().catch(console.error);
