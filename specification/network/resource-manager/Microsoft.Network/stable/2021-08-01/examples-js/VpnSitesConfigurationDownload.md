```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gives the sas-url to download the configurations for vpn-sites in a resource group.
 *
 * @summary Gives the sas-url to download the configurations for vpn-sites in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnSitesConfigurationDownload.json
 */
async function vpnSitesConfigurationDownload() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualWANName = "wan1";
  const request = {
    outputBlobSasUrl:
      "https://blobcortextesturl.blob.core.windows.net/folderforconfig/vpnFile?sp=rw&se=2018-01-10T03%3A42%3A04Z&sv=2017-04-17&sig=WvXrT5bDmDFfgHs%2Brz%2BjAu123eRCNE9BO0eQYcPDT7pY%3D&sr=b",
    vpnSites: ["/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/abc"],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnSitesConfiguration.beginDownloadAndWait(
    resourceGroupName,
    virtualWANName,
    request
  );
  console.log(result);
}

vpnSitesConfigurationDownload().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
