const { ACEProvisioningManagementPartnerAPI } = require("@azure/arm-managementpartner");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the management partner using the objectId and tenantId.
 *
 * @summary Get the management partner using the objectId and tenantId.
 * x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/GetPartnerDetailsNoPartnerId.json
 */
async function getPartnerDetails() {
  const credential = new DefaultAzureCredential();
  const client = new ACEProvisioningManagementPartnerAPI(credential);
  const result = await client.partners.get();
  console.log(result);
}

getPartnerDetails().catch(console.error);
