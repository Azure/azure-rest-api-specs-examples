const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAllServicesInSubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.services.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllServicesInSubscription().catch(console.error);
