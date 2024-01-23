using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;
using Azure.ResourceManager.Avs.Models;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Addons_Get_ArcReg.json
// this example is just showing the usage of "Addons_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudAddonResource created on azure
// for more information of creating AvsPrivateCloudAddonResource, please refer to the document of AvsPrivateCloudAddonResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string addonName = "arc";
ResourceIdentifier avsPrivateCloudAddonResourceId = AvsPrivateCloudAddonResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, addonName);
AvsPrivateCloudAddonResource avsPrivateCloudAddon = client.GetAvsPrivateCloudAddonResource(avsPrivateCloudAddonResourceId);

// invoke the operation
AvsPrivateCloudAddonResource result = await avsPrivateCloudAddon.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AvsPrivateCloudAddonData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
