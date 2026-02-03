using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VpnSitesConfigurationDownload.json
// this example is just showing the usage of "VpnSitesConfiguration_Download" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualWanResource created on azure
// for more information of creating VirtualWanResource, please refer to the document of VirtualWanResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualWanName = "wan1";
ResourceIdentifier virtualWanResourceId = VirtualWanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualWanName);
VirtualWanResource virtualWan = client.GetVirtualWanResource(virtualWanResourceId);

// invoke the operation
GetVpnSitesConfigurationContent content = new GetVpnSitesConfigurationContent(new Uri("https://blobcortextesturl.blob.core.windows.net/folderforconfig/vpnFile?sp=rw&se=2018-01-10T03%3A42%3A04Z&sv=2017-04-17&sig=WvXrT5bDmDFfgHs%2Brz%2BjAu123eRCNE9BO0eQYcPDT7pY%3D&sr=b"))
{
    VpnSites = { "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/abc" },
};
await virtualWan.DownloadVpnSitesConfigurationAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
