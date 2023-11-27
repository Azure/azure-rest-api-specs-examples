using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/GetOffer.json
// this example is just showing the usage of "Offers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublisherResource created on azure
// for more information of creating PublisherResource, please refer to the document of PublisherResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "myCluster";
string publisherName = "publisher1";
ResourceIdentifier publisherResourceId = PublisherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, publisherName);
PublisherResource publisher = client.GetPublisherResource(publisherResourceId);

// get the collection of this OfferResource
OfferCollection collection = publisher.GetOffers();

// invoke the operation
string offerName = "offer1";
NullableResponse<OfferResource> response = await collection.GetIfExistsAsync(offerName);
OfferResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    OfferData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
