using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/legacy/PrivateLinkResourcesListByService.json
// this example is just showing the usage of "PrivateLinkResources_ListByService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisServiceResource created on azure
// for more information of creating HealthcareApisServiceResource, please refer to the document of HealthcareApisServiceResource
string subscriptionId = "subid";
string resourceGroupName = "rgname";
string resourceName = "service1";
ResourceIdentifier healthcareApisServiceResourceId = HealthcareApisServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
HealthcareApisServiceResource healthcareApisService = client.GetHealthcareApisServiceResource(healthcareApisServiceResourceId);

// get the collection of this HealthcareApisServicePrivateLinkResource
HealthcareApisServicePrivateLinkResourceCollection collection = healthcareApisService.GetHealthcareApisServicePrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (HealthcareApisServicePrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthcareApisPrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
