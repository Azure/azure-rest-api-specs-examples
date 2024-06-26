const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or update a security application on the given subscription.
 *
 * @summary Creates or update a security application on the given subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/PutApplication_example.json
 */
async function createApplication() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const applicationId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
  const application = {
    description: "An application on critical recommendations",
    conditionSets: [
      {
        conditions: [{ operator: "contains", property: "$.Id", value: "-bil-" }],
      },
    ],
    displayName: "Admin's application",
    sourceResourceType: "Assessments",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.applicationOperations.createOrUpdate(applicationId, application);
  console.log(result);
}
