using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthcareApis;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/PrivateLinkResourceGet.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string groupName = "fhir";
NullableResponse<HealthcareApisServicePrivateLinkResource> response = await collection.GetIfExistsAsync(groupName);
HealthcareApisServicePrivateLinkResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthcareApisPrivateLinkResourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
