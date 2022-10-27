const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an ExpressRoute Circuit Authorization in a private cloud
 *
 * @summary Create or update an ExpressRoute Circuit Authorization in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/Authorizations_CreateOrUpdate.json
 */
async function authorizationsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const authorizationName = "authorization1";
  const authorization = {};
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.authorizations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateCloudName,
    authorizationName,
    authorization
  );
  console.log(result);
}

authorizationsCreateOrUpdate().catch(console.error);
