const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the certificates in the specified account.
 *
 * @summary Lists all of the certificates in the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/CertificateListWithFilter.json
 */
async function listCertificatesFilterAndSelect() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const select = "properties/format,properties/provisioningState";
  const filter =
    "properties/provisioningStateTransitionTime gt '2017-05-01' or properties/provisioningState eq 'Failed'";
  const options = {
    select,
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.certificateOperations.listByBatchAccount(
    resourceGroupName,
    accountName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCertificatesFilterAndSelect().catch(console.error);
