const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to checks whether a Managed Environment can be migrated to a target mode.
 *
 * @summary checks whether a Managed Environment can be migrated to a target mode.
 * x-ms-original-file: 2026-07-01/ManagedEnvironments_CheckMigrationEligibility_Ineligible.json
 */
async function checkMigrationEligibilityForAnIneligibleManagedEnvironment() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.checkMigrationEligibility(
    "examplerg",
    "my-environment",
    { targetMode: "Express" },
  );
  console.log(result);
}
