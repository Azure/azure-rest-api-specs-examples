```javascript
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managementpartner_2.0.1/sdk/managementpartner/arm-managementpartner/README.md) on how to add the SDK to your project and authenticate.
