const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a contact.
 *
 * @summary Creates a contact.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactCreate.json
 */
async function createAContact() {
  const subscriptionId =
    process.env["ORBITAL_SUBSCRIPTION_ID"] || "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
  const resourceGroupName = process.env["ORBITAL_RESOURCE_GROUP"] || "contoso-Rgp";
  const spacecraftName = "CONTOSO_SAT";
  const contactName = "contact1";
  const parameters = {
    contactProfile: {
      id: "/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP",
    },
    groundStationName: "EASTUS2_0",
    reservationEndTime: new Date("2023-02-22T11:10:45Z"),
    reservationStartTime: new Date("2023-02-22T10:58:30Z"),
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
