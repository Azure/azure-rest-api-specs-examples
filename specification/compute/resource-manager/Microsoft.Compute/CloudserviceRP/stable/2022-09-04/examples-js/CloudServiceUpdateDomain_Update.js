const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the role instances in the specified update domain.
 *
 * @summary Updates the role instances in the specified update domain.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceUpdateDomain_Update.json
 */
async function updateCloudServiceToSpecifiedDomain() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const updateDomain = 1;
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServicesUpdateDomain.beginWalkUpdateDomainAndWait(
    resourceGroupName,
    cloudServiceName,
    updateDomain,
  );
  console.log(result);
}
