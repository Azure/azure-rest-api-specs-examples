const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a role instance from a cloud service.
 *
 * @summary Gets a role instance from a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceRoleInstance_Get.json
 */
async function getCloudServiceRoleInstance() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const roleInstanceName = "{roleInstance-name}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServiceRoleInstances.get(
    roleInstanceName,
    resourceGroupName,
    cloudServiceName,
  );
  console.log(result);
}
