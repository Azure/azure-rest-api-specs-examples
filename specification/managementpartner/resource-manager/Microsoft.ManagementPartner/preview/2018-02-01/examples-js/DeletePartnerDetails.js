const { ACEProvisioningManagementPartnerAPI } = require("@azure/arm-managementpartner");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the management partner for the objectId and tenantId.
 *
 * @summary Delete the management partner for the objectId and tenantId.
 * x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/DeletePartnerDetails.json
 */
async function deletePartnerDetails() {
  const partnerId = "123456";
  const credential = new DefaultAzureCredential();
  const client = new ACEProvisioningManagementPartnerAPI(credential);
  const result = await client.partner.delete(partnerId);
  console.log(result);
}
