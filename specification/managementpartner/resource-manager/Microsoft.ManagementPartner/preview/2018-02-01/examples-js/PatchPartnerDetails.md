Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementpartner_2.0.1/sdk/managementpartner/arm-managementpartner/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ACEProvisioningManagementPartnerAPI } = require("@azure/arm-managementpartner");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the management partner for the objectId and tenantId.
 *
 * @summary Update the management partner for the objectId and tenantId.
 * x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/PatchPartnerDetails.json
 */
async function patchPartnerDetails() {
  const partnerId = "123456";
  const credential = new DefaultAzureCredential();
  const client = new ACEProvisioningManagementPartnerAPI(credential);
  const result = await client.partner.update(partnerId);
  console.log(result);
}

patchPartnerDetails().catch(console.error);
```
