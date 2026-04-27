const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get schema registry cluster by Id
 *
 * @summary get schema registry cluster by Id
 * x-ms-original-file: 2025-08-18-preview/Organization_GetSchemaRegistryClusterById_MaximumSet_Gen.json
 */
async function organizationGetSchemaRegistryClusterByIdMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "DC34558A-05D3-4370-AED8-75E60B381F94";
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.getSchemaRegistryClusterById(
    "rgconfluent",
    "hmhbrtw",
    "ztozszmpzhwevkpmaxslloijkicwt",
    "stfqijternpuzpleowkrbgzuutsgp",
  );
  console.log(result);
}
