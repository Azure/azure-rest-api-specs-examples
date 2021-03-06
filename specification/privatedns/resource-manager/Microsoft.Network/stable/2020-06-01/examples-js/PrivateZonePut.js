const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function putPrivateDnsZone() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const parameters = {
    location: "Global",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.privateZones.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateZoneName,
    parameters
  );
  console.log(result);
}

putPrivateDnsZone().catch(console.error);
