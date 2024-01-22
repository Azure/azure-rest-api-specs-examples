using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/CloudLinks_Delete.json
// this example is just showing the usage of "CloudLinks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsCloudLinkResource created on azure
// for more information of creating AvsCloudLinkResource, please refer to the document of AvsCloudLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string cloudLinkName = "cloudLink1";
ResourceIdentifier avsCloudLinkResourceId = AvsCloudLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, cloudLinkName);
AvsCloudLinkResource avsCloudLink = client.GetAvsCloudLinkResource(avsCloudLinkResourceId);

// invoke the operation
await avsCloudLink.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
