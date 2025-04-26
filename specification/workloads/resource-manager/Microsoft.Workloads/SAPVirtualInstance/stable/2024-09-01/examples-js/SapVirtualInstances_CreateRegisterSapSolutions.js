const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateRegisterSapSolutions.json
 */
async function registerExistingSAPSystemAsVirtualInstanceForSAPSolutions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapVirtualInstances.create("test-rg", "X00", {
    location: "northeurope",
    properties: {
      configuration: {
        centralServerVmId:
          "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0",
        configurationType: "Discovery",
      },
      environment: "NonProd",
      sapProduct: "S4HANA",
    },
    tags: { createdby: "abc@microsoft.com", test: "abc" },
  });
  console.log(result);
}
