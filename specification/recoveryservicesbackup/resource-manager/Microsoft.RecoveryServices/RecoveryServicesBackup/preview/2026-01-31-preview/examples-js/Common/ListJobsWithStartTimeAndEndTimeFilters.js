const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to provides a pageable list of jobs.
 *
 * @summary provides a pageable list of jobs.
 * x-ms-original-file: 2026-01-31-preview/Common/ListJobsWithStartTimeAndEndTimeFilters.json
 */
async function listJobsWithTimeFilter() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.backupJobs.list("NetSDKTestRsVault", "SwaggerTestRg", {
    filter: "startTime eq '2016-01-01 00:00:00 AM' and endTime eq '2017-11-29 00:00:00 AM'",
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
