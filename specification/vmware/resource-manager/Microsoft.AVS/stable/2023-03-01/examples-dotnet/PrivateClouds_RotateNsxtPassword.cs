using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;
using Azure.ResourceManager.Avs.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/PrivateClouds_RotateNsxtPassword.json
// this example is just showing the usage of "PrivateClouds_RotateNsxtPassword" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudResource created on azure
// for more information of creating AvsPrivateCloudResource, please refer to the document of AvsPrivateCloudResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
ResourceIdentifier avsPrivateCloudResourceId = AvsPrivateCloudResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName);
AvsPrivateCloudResource avsPrivateCloud = client.GetAvsPrivateCloudResource(avsPrivateCloudResourceId);

// invoke the operation
await avsPrivateCloud.RotateNsxtPasswordAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
