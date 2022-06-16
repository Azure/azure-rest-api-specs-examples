const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new certificate inside the specified account.
 *
 * @summary Creates a new certificate inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/CertificateCreate_MinimalCer.json
 */
async function createCertificateMinimalCer() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const certificateName = "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e";
  const parameters = {
    format: "Cer",
    data: "MIICrjCCAZagAwI...",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.certificateOperations.create(
    resourceGroupName,
    accountName,
    certificateName,
    parameters
  );
  console.log(result);
}

createCertificateMinimalCer().catch(console.error);
