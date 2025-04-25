const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to stops the SAP Application, that is the Application server instances and Central Services instance.
 *
 * @summary stops the SAP Application, that is the Application server instances and Central Services instance.
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_SoftStopVMAndSystem.json
 */
async function softStopTheVirtualMachineSAndTheSAPSystemOnIt() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapVirtualInstances.stop("test-rg", "X00", {
    body: { deallocateVm: true, softStopTimeoutSeconds: 300 },
  });
  console.log(result);
}
