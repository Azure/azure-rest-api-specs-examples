const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new user or updates an existing user's information on a Data Box Edge/Data Box Gateway device.
 *
 * @summary Creates a new user or updates an existing user's information on a Data Box Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/UserPut.json
 */
async function userPut() {
  const subscriptionId =
    process.env["DATABOXEDGE_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const name = "user1";
  const resourceGroupName = process.env["DATABOXEDGE_RESOURCE_GROUP"] || "GroupForEdgeAutomation";
  const user = {
    encryptedPassword: {
      encryptionAlgorithm: "None",
      encryptionCertThumbprint: "blah",
      value: "<value>",
    },
    shareAccessRights: [],
    userType: "Share",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.users.beginCreateOrUpdateAndWait(
    deviceName,
    name,
    resourceGroupName,
    user
  );
  console.log(result);
}
