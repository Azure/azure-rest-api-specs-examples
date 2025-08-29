const { ContainerServiceFleetClient } = require("@azure/arm-containerservicefleet");
const { DefaultAzureCredential } = require("@azure/identity");

async function createsAFleetResourceWithALongRunningOperation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceFleetClient(credential, subscriptionId);
  const result = await client.fleets.create("rg1", "fleet1", {
    tags: { tier: "production", archv2: "" },
    location: "East US",
    properties: {
      hubProfile: {
        dnsPrefix: "dnsprefix1",
        agentProfile: { vmSize: "Standard_DS1" },
      },
    },
  });
  console.log(result);
}
