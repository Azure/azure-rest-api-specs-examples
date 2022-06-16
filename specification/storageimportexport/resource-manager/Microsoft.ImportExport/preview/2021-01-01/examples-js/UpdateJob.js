const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateImportJob() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const jobName = "myJob";
  const resourceGroupName = "myResourceGroup";
  const body = {
    backupDriveManifest: true,
    logLevel: "Verbose",
    state: "",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const result = await client.jobs.update(jobName, resourceGroupName, body);
  console.log(result);
}

updateImportJob().catch(console.error);
