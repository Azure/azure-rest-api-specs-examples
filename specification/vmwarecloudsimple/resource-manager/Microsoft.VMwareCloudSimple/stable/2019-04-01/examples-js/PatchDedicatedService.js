const { VMwareCloudSimple } = require("@azure/arm-vmwarecloudsimple");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch dedicated cloud service's properties
 *
 * @summary Patch dedicated cloud service's properties
 * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/PatchDedicatedService.json
 */
async function patchDedicatedService() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const dedicatedCloudServiceName = "myService";
  const dedicatedCloudServiceRequest = {
    tags: { myTag: "tagValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new VMwareCloudSimple(credential, subscriptionId);
  const result = await client.dedicatedCloudServices.update(
    resourceGroupName,
    dedicatedCloudServiceName,
    dedicatedCloudServiceRequest
  );
  console.log(result);
}

patchDedicatedService().catch(console.error);
