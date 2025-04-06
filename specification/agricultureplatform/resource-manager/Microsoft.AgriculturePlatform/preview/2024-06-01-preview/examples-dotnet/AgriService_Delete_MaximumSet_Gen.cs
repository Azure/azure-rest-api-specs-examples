using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AgriculturePlatform.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.AgriculturePlatform;

// Generated from example definition: 2024-06-01-preview/AgriService_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "AgriServiceResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AgricultureServiceResource created on azure
// for more information of creating AgricultureServiceResource, please refer to the document of AgricultureServiceResource
string subscriptionId = "83D293F5-DEFD-4D48-B120-1DC713BE338A";
string resourceGroupName = "rgopenapi";
string agriServiceResourceName = "abc123";
ResourceIdentifier agricultureServiceResourceId = AgricultureServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, agriServiceResourceName);
AgricultureServiceResource agricultureService = client.GetAgricultureServiceResource(agricultureServiceResourceId);

// invoke the operation
await agricultureService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
