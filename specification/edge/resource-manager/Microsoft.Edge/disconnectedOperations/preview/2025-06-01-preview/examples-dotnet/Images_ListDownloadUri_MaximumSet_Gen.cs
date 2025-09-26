using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DisconnectedOperations.Models;
using Azure.ResourceManager.DisconnectedOperations;

// Generated from example definition: 2025-06-01-preview/Images_ListDownloadUri_MaximumSet_Gen.json
// this example is just showing the usage of "Images_ListDownloadUri" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DisconnectedOperationsImageResource created on azure
// for more information of creating DisconnectedOperationsImageResource, please refer to the document of DisconnectedOperationsImageResource
string subscriptionId = "51DB5DE7-A66C-4789-BFFF-9F75C95A0201";
string resourceGroupName = "rgdisconnectedOperations";
string name = "g_-5-160";
string imageName = "1Q6lGV4V65j-1";
ResourceIdentifier disconnectedOperationsImageResourceId = DisconnectedOperationsImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, imageName);
DisconnectedOperationsImageResource disconnectedOperationsImage = client.GetDisconnectedOperationsImageResource(disconnectedOperationsImageResourceId);

// invoke the operation
DisconnectedOperationsImageDownloadResult result = await disconnectedOperationsImage.GetDownloadUriAsync();

Console.WriteLine($"Succeeded: {result}");
