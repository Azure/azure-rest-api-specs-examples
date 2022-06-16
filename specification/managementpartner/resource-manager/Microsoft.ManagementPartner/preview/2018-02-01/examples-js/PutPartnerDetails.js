const { ACEProvisioningManagementPartnerAPI } = require("@azure/arm-managementpartner");
const { DefaultAzureCredential } = require("@azure/identity");

async function putPartnerDetails() {
  const partnerId = "123456";
  const credential = new DefaultAzureCredential();
  const client = new ACEProvisioningManagementPartnerAPI(credential);
  const result = await client.partner.create(partnerId);
  console.log(result);
}

putPartnerDetails().catch(console.error);
