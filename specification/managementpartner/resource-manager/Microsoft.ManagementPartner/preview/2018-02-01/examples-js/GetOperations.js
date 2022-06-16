const { ACEProvisioningManagementPartnerAPI } = require("@azure/arm-managementpartner");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the operations.
 *
 * @summary List all the operations.
 * x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/GetOperations.json
 */
async function getOperations() {
  const credential = new DefaultAzureCredential();
  const client = new ACEProvisioningManagementPartnerAPI(credential);
  const resArray = new Array();
  for await (let item of client.operation.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getOperations().catch(console.error);
