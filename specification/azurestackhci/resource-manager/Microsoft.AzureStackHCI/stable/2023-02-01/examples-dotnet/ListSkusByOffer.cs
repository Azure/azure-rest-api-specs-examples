using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/ListSkusByOffer.json
// this example is just showing the usage of "Skus_ListByOffer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OfferResource created on azure
// for more information of creating OfferResource, please refer to the document of OfferResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "myCluster";
string publisherName = "publisher1";
string offerName = "offer1";
ResourceIdentifier offerResourceId = OfferResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, publisherName, offerName);
OfferResource offer = client.GetOfferResource(offerResourceId);

// get the collection of this HciSkuResource
HciSkuCollection collection = offer.GetHciSkus();

// invoke the operation and iterate over the result
await foreach (HciSkuResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HciSkuData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
