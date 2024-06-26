const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns dedicated cloud node by its name
 *
 * @summary Returns dedicated cloud node by its name
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateDedicatedCloudNode.json
 */
async function createDedicatedCloudNode() {
  const subscriptionId = process.env["VMWARECLOUDSIMPLE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["VMWARECLOUDSIMPLE_RESOURCE_GROUP"] || "myResourceGroup";
  const referer = "https://management.azure.com/";
  const dedicatedCloudNodeName = "myNode";
  const dedicatedCloudNodeRequest = {
    namePropertiesSkuDescriptionName: "CS28-Node",
    availabilityZoneId: "az1",
    idPropertiesSkuDescriptionId: "general",
    location: "westus",
    nodesCount: 1,
    placementGroupId: "n1",
    purchaseId: "56acbd46-3d36-4bbf-9b08-57c30fdf6932",
    sku: { name: "VMware_CloudSimple_CS28" },
  };
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.dedicatedCloudNodes.beginCreateOrUpdateAndWait(
    resourceGroupName,
    referer,
    dedicatedCloudNodeName,
    dedicatedCloudNodeRequest
  );
  console.log(result);
}
