const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing job.
 *
 * @summary Updates the properties of an existing job.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsPatchCmk.json
 */
async function jobsPatchCmk() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "SdkRg7937";
  const jobName = "SdkJob1735";
  const jobResourceUpdateParameter = {
    details: {
      keyEncryptionKey: {
        kekType: "CustomerManaged",
        kekUrl: "https://sdkkeyvault.vault.azure.net/keys/SSDKEY/",
        kekVaultResourceID:
          "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.KeyVault/vaults/SDKKeyVault",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.beginUpdateAndWait(
    resourceGroupName,
    jobName,
    jobResourceUpdateParameter
  );
  console.log(result);
}

jobsPatchCmk().catch(console.error);
