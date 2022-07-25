const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This API provides configuration details specific to given region/location at Resource group level.
 *
 * @summary This API provides configuration details specific to given region/location at Resource group level.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/RegionConfigurationByResourceGroup.json
 */
async function regionConfigurationByResourceGroup() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "SdkRg4981";
  const location = "westus";
  const regionConfigurationRequest = {
    scheduleAvailabilityRequest: {
      skuName: "DataBox",
      storageLocation: "westus",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.service.regionConfigurationByResourceGroup(
    resourceGroupName,
    location,
    regionConfigurationRequest
  );
  console.log(result);
}

regionConfigurationByResourceGroup().catch(console.error);
