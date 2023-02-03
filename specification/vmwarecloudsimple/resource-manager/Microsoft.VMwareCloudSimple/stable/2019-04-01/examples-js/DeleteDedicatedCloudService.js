const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete dedicate cloud service
 *
 * @summary Delete dedicate cloud service
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/DeleteDedicatedCloudService.json
 */
async function deleteDedicatedCloudService() {
  const subscriptionId = process.env["VMWARECLOUDSIMPLE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["VMWARECLOUDSIMPLE_RESOURCE_GROUP"] || "myResourceGroup";
  const dedicatedCloudServiceName = "myService";
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.dedicatedCloudServices.beginDeleteAndWait(
    resourceGroupName,
    dedicatedCloudServiceName
  );
  console.log(result);
}
