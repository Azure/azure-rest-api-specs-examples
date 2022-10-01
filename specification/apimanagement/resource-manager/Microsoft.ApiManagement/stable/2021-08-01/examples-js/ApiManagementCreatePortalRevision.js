const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new developer portal's revision by running the portal's publishing. The `isCurrent` property indicates if the revision is publicly accessible.
 *
 * @summary Creates a new developer portal's revision by running the portal's publishing. The `isCurrent` property indicates if the revision is publicly accessible.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreatePortalRevision.json
 */
async function apiManagementCreatePortalRevision() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const portalRevisionId = "20201112101010";
  const parameters = {
    description: "portal revision 1",
    isCurrent: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.portalRevision.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    portalRevisionId,
    parameters
  );
  console.log(result);
}

apiManagementCreatePortalRevision().catch(console.error);
