const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the details of the API specified by its identifier in the format specified to the Storage Blob with SAS Key valid for 5 minutes.
 *
 * @summary Gets the details of the API specified by its identifier in the format specified to the Storage Blob with SAS Key valid for 5 minutes.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspaceApiExportInOpenApi3dot0.json
 */
async function apiManagementGetWorkspaceApiExportInOpenApi3Dot0() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const apiId = "aid9676";
  const format = "openapi-link";
  const exportParam = "true";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceApiExport.get(
    resourceGroupName,
    serviceName,
    workspaceId,
    apiId,
    format,
    exportParam,
  );
  console.log(result);
}
