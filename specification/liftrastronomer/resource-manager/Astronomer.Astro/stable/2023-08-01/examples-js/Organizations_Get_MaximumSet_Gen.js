const { AstroManagementClient } = require("@azure/arm-astro");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a OrganizationResource
 *
 * @summary Get a OrganizationResource
 * x-ms-original-file: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_Get_MaximumSet_Gen.json
 */
async function organizationsGetGeneratedByMaximumSetRule() {
  const subscriptionId =
    process.env["ASTRO_SUBSCRIPTION_ID"] || "43454B17-172A-40FE-80FA-549EA23D12B3";
  const resourceGroupName = process.env["ASTRO_RESOURCE_GROUP"] || "rgastronomer";
  const organizationName = "S PS";
  const credential = new DefaultAzureCredential();
  const client = new AstroManagementClient(credential, subscriptionId);
  const result = await client.organizations.get(resourceGroupName, organizationName);
  console.log(result);
}
