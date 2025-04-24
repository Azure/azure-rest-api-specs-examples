using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PineconeVectorDB.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PineconeVectorDB;

// Generated from example definition: 2024-10-22-preview/Organizations_ListByResourceGroup_MaximumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "76a38ef6-c8c1-4f0d-bfe0-00ec782c8077";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PineconeVectorDBOrganizationResource
PineconeVectorDBOrganizationCollection collection = resourceGroupResource.GetPineconeVectorDBOrganizations();

// invoke the operation and iterate over the result
await foreach (PineconeVectorDBOrganizationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PineconeVectorDBOrganizationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
