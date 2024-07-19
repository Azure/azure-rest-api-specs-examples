using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/OriginGroups_Delete.json
// this example is just showing the usage of "CdnOriginGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CdnOriginGroupResource created on azure
// for more information of creating CdnOriginGroupResource, please refer to the document of CdnOriginGroupResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string endpointName = "endpoint1";
string originGroupName = "originGroup1";
ResourceIdentifier cdnOriginGroupResourceId = CdnOriginGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointName, originGroupName);
CdnOriginGroupResource cdnOriginGroup = client.GetCdnOriginGroupResource(cdnOriginGroupResourceId);

// invoke the operation
await cdnOriginGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
