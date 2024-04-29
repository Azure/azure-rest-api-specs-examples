const { GuestConfigurationClient } = require("@azure/arm-guestconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an association between a ARC machine and guest configuration
 *
 * @summary Creates an association between a ARC machine and guest configuration
 * x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/createOrUpdateGuestConfigurationHCRPAssignment.json
 */
async function createOrUpdateGuestConfigurationAssignment() {
  const subscriptionId = process.env["GUESTCONFIGURATION_SUBSCRIPTION_ID"] || "mySubscriptionId";
  const guestConfigurationAssignmentName = "NotInstalledApplicationForWindows";
  const resourceGroupName =
    process.env["GUESTCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroupName";
  const machineName = "myMachineName";
  const parameters = {
    name: "NotInstalledApplicationForWindows",
    location: "westcentralus",
    properties: {
      context: "Azure policy",
      guestConfiguration: {
        name: "NotInstalledApplicationForWindows",
        assignmentType: "ApplyAndAutoCorrect",
        configurationParameter: [
          {
            name: "[InstalledApplication]NotInstalledApplicationResource1;Name",
            value: "NotePad,sql",
          },
        ],
        contentHash: "123contenthash",
        contentUri: "https://thisisfake/pacakge",
        version: "1.0.0.3",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new GuestConfigurationClient(credential, subscriptionId);
  const result = await client.guestConfigurationHcrpAssignments.createOrUpdate(
    guestConfigurationAssignmentName,
    resourceGroupName,
    machineName,
    parameters,
  );
  console.log(result);
}
