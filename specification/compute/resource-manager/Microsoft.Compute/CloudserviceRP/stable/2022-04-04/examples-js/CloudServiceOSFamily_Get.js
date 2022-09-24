const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a guest operating system family that can be specified in the XML service configuration (.cscfg) for a cloud service.
 *
 * @summary Gets properties of a guest operating system family that can be specified in the XML service configuration (.cscfg) for a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudServiceOSFamily_Get.json
 */
async function getCloudServiceOSFamily() {
  const subscriptionId = "{subscription-id}";
  const location = "westus2";
  const osFamilyName = "3";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServiceOperatingSystems.getOSFamily(location, osFamilyName);
  console.log(result);
}

getCloudServiceOSFamily().catch(console.error);
