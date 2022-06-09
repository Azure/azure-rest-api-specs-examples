```javascript
const { ACEProvisioningManagementPartnerAPI } = require("@azure/arm-managementpartner");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the management partner using the partnerId, objectId and tenantId.
 *
 * @summary Get the management partner using the partnerId, objectId and tenantId.
 * x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/GetPartnerDetails.json
 */
async function getPartnerDetails() {
  const partnerId = "123456";
  const credential = new DefaultAzureCredential();
  const client = new ACEProvisioningManagementPartnerAPI(credential);
  const result = await client.partner.get(partnerId);
  console.log(result);
}

getPartnerDetails().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementpartner_2.0.1/sdk/managementpartner/arm-managementpartner/README.md) on how to add the SDK to your project and authenticate.
