const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the availability of name for resource
 *
 * @summary Check the availability of name for resource
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/CheckNameAvailability.json
 */
async function nameAvailability() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const nameAvailabilityRequest = {
    name: "name1",
    type: "Microsoft.DBforMySQL",
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.checkNameAvailability.execute(nameAvailabilityRequest);
  console.log(result);
}

nameAvailability().catch(console.error);
