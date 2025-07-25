const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Gets status of a web app backup that may be in progress, including secrets associated with the backup, such as the Azure Storage SAS URL. Also can be used to update the SAS URL for the backup if a new URL is passed in the request body.
 *
 * @summary Description for Gets status of a web app backup that may be in progress, including secrets associated with the backup, such as the Azure Storage SAS URL. Also can be used to update the SAS URL for the backup if a new URL is passed in the request body.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetWebAppBackupWithSecrets.json
 */
async function getWebAppBackupWithSecrets() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "sitef6141";
  const backupId = "12345";
  const request = {
    backupName: "abcdwe",
    backupSchedule: {
      frequencyInterval: 7,
      frequencyUnit: "Day",
      keepAtLeastOneBackup: true,
      retentionPeriodInDays: 30,
      startTime: new Date("2022-09-02T17:33:11.641Z"),
    },
    databases: [
      {
        name: "backenddb",
        connectionString:
          "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]",
        connectionStringName: "backend",
        databaseType: "SqlAzure",
      },
      {
        name: "statsdb",
        connectionString:
          "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]",
        connectionStringName: "stats",
        databaseType: "SqlAzure",
      },
    ],
    enabled: true,
    storageAccountUrl:
      "DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.listBackupStatusSecrets(
    resourceGroupName,
    name,
    backupId,
    request,
  );
  console.log(result);
}
