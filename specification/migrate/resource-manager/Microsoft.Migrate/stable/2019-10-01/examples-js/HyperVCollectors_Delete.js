const { AzureMigrateV2 } = require("@azure/arm-migrate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Hyper-V collector from the project.
 *
 * @summary Delete a Hyper-V collector from the project.
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/HyperVCollectors_Delete.json
 */
async function hyperVCollectorsDelete() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "8c3c936a-c09b-4de3-830b-3f5f244d72e9";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "contosoithyperv";
  const projectName = "migrateprojectce73project";
  const hyperVCollectorName = "migrateprojectce73collector";
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateV2(credential, subscriptionId);
  const result = await client.hyperVCollectors.delete(
    resourceGroupName,
    projectName,
    hyperVCollectorName
  );
  console.log(result);
}
