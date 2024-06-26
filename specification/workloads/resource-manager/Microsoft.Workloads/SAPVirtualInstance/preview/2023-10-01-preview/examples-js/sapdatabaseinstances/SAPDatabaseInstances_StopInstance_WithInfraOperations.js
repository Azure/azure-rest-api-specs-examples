const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops the database instance of the SAP system.
 *
 * @summary Stops the database instance of the SAP system.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapdatabaseinstances/SAPDatabaseInstances_StopInstance_WithInfraOperations.json
 */
async function stopTheDatabaseInstanceOfTheSapSystemAndTheUnderlyingVirtualMachineS() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const databaseInstanceName = "db0";
  const body = { deallocateVm: true, softStopTimeoutSeconds: 0 };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPDatabaseInstances.beginStopInstanceAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
    databaseInstanceName,
    options,
  );
  console.log(result);
}
