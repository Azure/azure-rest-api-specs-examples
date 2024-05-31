const { InformaticaDataManagement } = require("@azure/arm-informaticadatamanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a InformaticaOrganizationResource
 *
 * @summary Create a InformaticaOrganizationResource
 * x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_CreateOrUpdate_MinimumSet_Gen.json
 */
async function organizationsCreateOrUpdateMin() {
  const subscriptionId =
    process.env["INFORMATICA_SUBSCRIPTION_ID"] || "3599DA28-E346-4D9F-811E-189C0445F0FE";
  const resourceGroupName = process.env["INFORMATICA_RESOURCE_GROUP"] || "rgopenapi";
  const organizationName = "n6v";
  const resource = {
    location: "pamjoudtssthlbhrnfjidr",
  };
  const credential = new DefaultAzureCredential();
  const client = new InformaticaDataManagement(credential, subscriptionId);
  const result = await client.organizations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    organizationName,
    resource,
  );
  console.log(result);
}
