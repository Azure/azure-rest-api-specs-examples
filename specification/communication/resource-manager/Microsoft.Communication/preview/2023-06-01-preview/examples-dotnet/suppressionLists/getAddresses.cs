using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/preview/2023-06-01-preview/examples/suppressionLists/getAddresses.json
// this example is just showing the usage of "SuppressionListAddresses_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SuppressionListResource created on azure
// for more information of creating SuppressionListResource, please refer to the document of SuppressionListResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "contosoResourceGroup";
string emailServiceName = "contosoEmailService";
string domainName = "contoso.com";
string suppressionListName = "aaaa1111-bbbb-2222-3333-aaaa11112222";
ResourceIdentifier suppressionListResourceId = SuppressionListResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, emailServiceName, domainName, suppressionListName);
SuppressionListResource suppressionListResource = client.GetSuppressionListResource(suppressionListResourceId);

// get the collection of this SuppressionListAddressResource
SuppressionListAddressResourceCollection collection = suppressionListResource.GetSuppressionListAddressResources();

// invoke the operation and iterate over the result
await foreach (SuppressionListAddressResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SuppressionListAddressResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
