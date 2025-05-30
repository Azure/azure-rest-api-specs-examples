const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This API provides configuration details specific to given region/location at Resource group level.
 *
 * @summary This API provides configuration details specific to given region/location at Resource group level.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/RegionConfigurationByResourceGroup.json
 */
async function regionConfigurationByResourceGroup() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const location = "westus";
  const regionConfigurationRequest = {
    deviceCapabilityRequest: { model: "DataBoxDisk", skuName: "DataBoxDisk" },
    scheduleAvailabilityRequest: {
      model: "DataBox",
      skuName: "DataBox",
      storageLocation: "westus",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.service.regionConfigurationByResourceGroup(
    resourceGroupName,
    location,
    regionConfigurationRequest,
  );
  console.log(result);
}
