using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MongoDBAtlas.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MongoDBAtlas;

// Generated from example definition: 2024-11-18-preview/Organizations_ListByResourceGroup_MinimumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "422A4D59-A5BC-4DBB-8831-EC666633F64F";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MongoDBAtlasOrganizationResource
MongoDBAtlasOrganizationCollection collection = resourceGroupResource.GetMongoDBAtlasOrganizations();

// invoke the operation and iterate over the result
await foreach (MongoDBAtlasOrganizationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MongoDBAtlasOrganizationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
