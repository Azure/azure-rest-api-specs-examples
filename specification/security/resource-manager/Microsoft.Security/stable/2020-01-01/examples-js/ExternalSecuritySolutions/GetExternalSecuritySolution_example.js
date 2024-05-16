const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a specific external Security Solution.
 *
 * @summary Gets a specific external Security Solution.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ExternalSecuritySolutions/GetExternalSecuritySolution_example.json
 */
async function getExternalSecuritySolution() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "defaultresourcegroup-eus";
  const ascLocation = "centralus";
  const externalSecuritySolutionsName =
    "aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.externalSecuritySolutions.get(
    resourceGroupName,
    ascLocation,
    externalSecuritySolutionsName,
  );
  console.log(result);
}
