using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthDataAIServices.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HealthDataAIServices;

// Generated from example definition: 2024-09-20/PrivateLinks_ListByDeidService_MaximumSet_Gen.json
// this example is just showing the usage of "PrivateLinkResource_ListByDeidService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeidServiceResource created on azure
// for more information of creating DeidServiceResource, please refer to the document of DeidServiceResource
string subscriptionId = "F21BB31B-C214-42C0-ACF0-DACCA05D3011";
string resourceGroupName = "rgopenapi";
string deidServiceName = "deidTest";
ResourceIdentifier deidServiceResourceId = DeidServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deidServiceName);
DeidServiceResource deidService = client.GetDeidServiceResource(deidServiceResourceId);

// invoke the operation and iterate over the result
await foreach (HealthDataAIServicesPrivateLinkResourceData item in deidService.GetPrivateLinksAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
