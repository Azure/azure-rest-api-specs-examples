const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a cloud service.
 *
 * @summary Update a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudService_Update_ToIncludeTags.json
 */
async function updateExistingCloudServiceToAddTags() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const parameters = { tags: { documentation: "RestAPI" } };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServices.beginUpdateAndWait(
    resourceGroupName,
    cloudServiceName,
    options,
  );
  console.log(result);
}
