const { DataCatalogRestClient } = require("@azure/arm-datacatalog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Update Azure Data Catalog Service operation can be used to update the existing deployment. The update call only supports the properties listed in the PATCH body.
 *
 * @summary The Update Azure Data Catalog Service operation can be used to update the existing deployment. The update call only supports the properties listed in the PATCH body.
 * x-ms-original-file: specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/UpdateADCCatalog.json
 */
async function updateAzureDataCatalogService() {
  const subscriptionId = "12345678-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const catalogName = "exampleCatalog";
  const properties = {
    admins: [
      {
        objectId: "99999999-9999-9999-999999999999",
        upn: "myupn@microsoft.com",
      },
    ],
    enableAutomaticUnitAdjustment: false,
    location: "North US",
    sku: "Standard",
    tags: { mykey: "myvalue", mykey2: "myvalue2" },
    units: 1,
    users: [
      {
        objectId: "99999999-9999-9999-999999999999",
        upn: "myupn@microsoft.com",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DataCatalogRestClient(credential, subscriptionId);
  const result = await client.aDCCatalogs.update(resourceGroupName, catalogName, properties);
  console.log(result);
}

updateAzureDataCatalogService().catch(console.error);
