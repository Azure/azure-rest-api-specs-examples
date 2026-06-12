const { SqlVirtualMachineManagementClient } = require("@azure/arm-sqlvirtualmachine");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a SQL virtual machine.
 *
 * @summary creates or updates a SQL virtual machine.
 * x-ms-original-file: 2023-10-01/CreateOrUpdateVirtualMachineWithVMGroup.json
 */
async function createsOrUpdatesASQLVirtualMachineAndJoinsItToASQLVirtualMachineGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlVirtualMachineManagementClient(credential, subscriptionId);
  const result = await client.sqlVirtualMachines.createOrUpdate("testrg", "testvm", {
    location: "northeurope",
    sqlVirtualMachineGroupResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup",
    virtualMachineResourceId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm2",
    wsfcDomainCredentials: {
      clusterBootstrapAccountPassword: "<Password>",
      clusterOperatorAccountPassword: "<Password>",
      sqlServiceAccountPassword: "<Password>",
    },
    wsfcStaticIp: "10.0.0.7",
  });
  console.log(result);
}
