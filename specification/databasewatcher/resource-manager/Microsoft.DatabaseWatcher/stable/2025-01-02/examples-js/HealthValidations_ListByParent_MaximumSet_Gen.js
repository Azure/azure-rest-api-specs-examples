const { DatabaseWatcherClient } = require("@azure/arm-databasewatcher");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list HealthValidation resources by Watcher
 *
 * @summary list HealthValidation resources by Watcher
 * x-ms-original-file: 2025-01-02/HealthValidations_ListByParent_MaximumSet_Gen.json
 */
async function healthValidationsListByParentMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "A76F9850-996B-40B3-94D4-C98110A0EEC9";
  const client = new DatabaseWatcherClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.healthValidations.listByParent("rgWatcher", "testWatcher")) {
    resArray.push(item);
  }

  console.log(resArray);
}
