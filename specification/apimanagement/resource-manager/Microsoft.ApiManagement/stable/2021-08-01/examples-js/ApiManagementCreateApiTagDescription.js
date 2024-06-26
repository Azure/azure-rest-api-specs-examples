const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create/Update tag description in scope of the Api.
 *
 * @summary Create/Update tag description in scope of the Api.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiTagDescription.json
 */
async function apiManagementCreateApiTagDescription() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "5931a75ae4bbd512a88c680b";
  const tagDescriptionId = "tagId1";
  const parameters = {
    description:
      "Some description that will be displayed for operation's tag if the tag is assigned to operation of the API",
    externalDocsDescription: "Description of the external docs resource",
    externalDocsUrl: "http://some.url/additionaldoc",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiTagDescription.createOrUpdate(
    resourceGroupName,
    serviceName,
    apiId,
    tagDescriptionId,
    parameters
  );
  console.log(result);
}
