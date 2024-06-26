const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts the SAP Application Server Instance.
 *
 * @summary Starts the SAP Application Server Instance.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapapplicationinstances/SAPApplicationServerInstances_StartInstance_WithInfraOperations.json
 */
async function startVirtualMachineAndTheSapApplicationServerInstanceOnIt() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const applicationInstanceName = "app01";
  const body = { startVm: true };
  const options = {
    body,
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPApplicationServerInstances.beginStartInstanceAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
    applicationInstanceName,
    options,
  );
  console.log(result);
}
