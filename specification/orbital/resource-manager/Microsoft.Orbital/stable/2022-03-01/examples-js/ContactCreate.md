```javascript
const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a contact.
 *
 * @summary Creates a contact.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactCreate.json
 */
async function createAContact() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const spacecraftName = "AQUA";
  const contactName = "contact1";
  const parameters = {
    contactProfile: {
      id: "/subscriptions/subId/resourceGroups/rg/Microsoft.Orbital/contactProfiles/AQUA_DIRECTPLAYBACK_WITH_UPLINK",
    },
    groundStationName: "westus_gs1",
    reservationEndTime: new Date("2020-07-16T20:55:00.00Z"),
    reservationStartTime: new Date("2020-07-16T20:35:00.00Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.contacts.beginCreateAndWait(
    resourceGroupName,
    spacecraftName,
    contactName,
    parameters
  );
  console.log(result);
}

createAContact().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-orbital_1.0.0/sdk/orbital/arm-orbital/README.md) on how to add the SDK to your project and authenticate.
