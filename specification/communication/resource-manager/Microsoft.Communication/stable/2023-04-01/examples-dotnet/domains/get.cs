using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Communication.Models;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/stable/2023-04-01/examples/domains/get.json
// this example is just showing the usage of "Domains_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EmailServiceResource created on azure
// for more information of creating EmailServiceResource, please refer to the document of EmailServiceResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "MyResourceGroup";
string emailServiceName = "MyEmailServiceResource";
ResourceIdentifier emailServiceResourceId = EmailServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, emailServiceName);
EmailServiceResource emailServiceResource = client.GetEmailServiceResource(emailServiceResourceId);

// get the collection of this CommunicationDomainResource
CommunicationDomainResourceCollection collection = emailServiceResource.GetCommunicationDomainResources();

// invoke the operation
string domainName = "mydomain.com";
NullableResponse<CommunicationDomainResource> response = await collection.GetIfExistsAsync(domainName);
CommunicationDomainResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CommunicationDomainResourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
