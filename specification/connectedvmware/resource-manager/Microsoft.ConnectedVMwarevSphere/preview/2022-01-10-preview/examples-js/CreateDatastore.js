const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update datastore.
 *
 * @summary Create Or Update datastore.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/CreateDatastore.json
 */
async function createDatastore() {
  const subscriptionId =
    process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "testrg";
  const datastoreName = "HRDatastore";
  const body = {
    extendedLocation: {
      name: "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso",
      type: "customLocation",
    },
    location: "East US",
    moRefId: "aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
    vCenterId:
      "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter",
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.datastores.beginCreateAndWait(
    resourceGroupName,
    datastoreName,
    options
  );
  console.log(result);
}
