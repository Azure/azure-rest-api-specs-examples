const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Data Plane access.
 *
 * @summary Get Data Plane access.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetDataPlaneAccess.json
 */
async function factoriesGetDataPlaneAccess() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const policy = {
    accessResourcePath: "",
    expireTime: "2018-11-10T09:46:20.2659347Z",
    permissions: "r",
    profileName: "DefaultProfile",
    startTime: "2018-11-10T02:46:20.2659347Z",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.factories.getDataPlaneAccess(resourceGroupName, factoryName, policy);
  console.log(result);
}

factoriesGetDataPlaneAccess().catch(console.error);
