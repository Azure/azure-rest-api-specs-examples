const { InformaticaDataManagement } = require("@azure/arm-informaticadatamanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a InformaticaServerlessRuntimeResource
 *
 * @summary Get a InformaticaServerlessRuntimeResource
 * x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Get_MaximumSet_Gen.json
 */
async function serverlessRuntimesGet() {
  const subscriptionId =
    process.env["INFORMATICA_SUBSCRIPTION_ID"] || "3599DA28-E346-4D9F-811E-189C0445F0FE";
  const resourceGroupName = process.env["INFORMATICA_RESOURCE_GROUP"] || "rgopenapi";
  const organizationName = "e3Y";
  const serverlessRuntimeName = "48-";
  const credential = new DefaultAzureCredential();
  const client = new InformaticaDataManagement(credential, subscriptionId);
  const result = await client.serverlessRuntimes.get(
    resourceGroupName,
    organizationName,
    serverlessRuntimeName,
  );
  console.log(result);
}
