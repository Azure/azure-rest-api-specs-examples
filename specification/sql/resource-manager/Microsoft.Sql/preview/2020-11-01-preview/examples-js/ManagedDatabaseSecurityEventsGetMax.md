Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of security events.
 *
 * @summary Gets a list of security events.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSecurityEventsGetMax.json
 */
async function getTheManagedDatabaseSecurityEventsWithMaximalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg";
  const managedInstanceName = "testcl";
  const databaseName = "database1";
  const filter = "ShowServerRecords eq true";
  const skip = 0;
  const top = 1;
  const skiptoken =
    "eyJCbG9iTmFtZURhdGVUaW1lIjoiXC9EYXRlKDE1MTIyODg4MTIwMTArMDIwMClcLyIsIkJsb2JOYW1lUm9sbG92ZXJJbmRleCI6IjAiLCJFbmREYXRlIjoiXC9EYXRlKDE1MTI0NjYyMDA1MjkpXC8iLCJJc1NraXBUb2tlblNldCI6ZmFsc2UsIklzVjJCbG9iVGltZUZvcm1hdCI6dHJ1ZSwiU2hvd1NlcnZlclJlY29yZHMiOmZhbHNlLCJTa2lwVmFsdWUiOjAsIlRha2VWYWx1ZSI6MTB9";
  const options = {
    filter,
    skip,
    top,
    skiptoken,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managedDatabaseSecurityEvents.listByDatabase(
    resourceGroupName,
    managedInstanceName,
    databaseName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getTheManagedDatabaseSecurityEventsWithMaximalParameters().catch(console.error);
```
