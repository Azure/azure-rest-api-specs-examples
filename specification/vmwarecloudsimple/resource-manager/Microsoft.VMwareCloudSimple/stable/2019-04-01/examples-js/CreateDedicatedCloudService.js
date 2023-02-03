const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create dedicate cloud service
 *
 * @summary Create dedicate cloud service
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateDedicatedCloudService.json
 */
async function createDedicatedCloudService() {
  const subscriptionId = process.env["VMWARECLOUDSIMPLE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["VMWARECLOUDSIMPLE_RESOURCE_GROUP"] || "myResourceGroup";
  const dedicatedCloudServiceName = "myService";
  const dedicatedCloudServiceRequest = {
    gatewaySubnet: "10.0.0.0",
    location: "westus",
  };
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.dedicatedCloudServices.createOrUpdate(
    resourceGroupName,
    dedicatedCloudServiceName,
    dedicatedCloudServiceRequest
  );
  console.log(result);
}
