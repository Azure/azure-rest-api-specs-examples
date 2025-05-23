using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileServicesGetUsage.json
// this example is just showing the usage of "FileServices_GetServiceUsage" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FileServiceUsageResource created on azure
// for more information of creating FileServiceUsageResource, please refer to the document of FileServiceUsageResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "res4410";
string accountName = "sto8607";
ResourceIdentifier fileServiceUsageResourceId = FileServiceUsageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
FileServiceUsageResource fileServiceUsage = client.GetFileServiceUsageResource(fileServiceUsageResourceId);

// invoke the operation
FileServiceUsageResource result = await fileServiceUsage.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FileServiceUsageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
