using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Communication;
using Azure.ResourceManager.Communication.Models;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/domains/createOrUpdate.json
// this example is just showing the usage of "Domains_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
CommunicationDomainResourceData data = new CommunicationDomainResourceData(new AzureLocation("Global"))
{
    DomainManagement = DomainManagement.CustomerManaged,
};
ArmOperation<CommunicationDomainResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, domainName, data);
CommunicationDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CommunicationDomainResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
