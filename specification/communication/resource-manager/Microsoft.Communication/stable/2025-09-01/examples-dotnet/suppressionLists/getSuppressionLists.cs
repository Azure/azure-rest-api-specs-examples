using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/getSuppressionLists.json
// this example is just showing the usage of "SuppressionLists_ListByDomain" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CommunicationDomainResource created on azure
// for more information of creating CommunicationDomainResource, please refer to the document of CommunicationDomainResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "contosoResourceGroup";
string emailServiceName = "contosoEmailService";
string domainName = "contoso.com";
ResourceIdentifier communicationDomainResourceId = CommunicationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, emailServiceName, domainName);
CommunicationDomainResource communicationDomainResource = client.GetCommunicationDomainResource(communicationDomainResourceId);

// get the collection of this EmailSuppressionListResource
EmailSuppressionListCollection collection = communicationDomainResource.GetEmailSuppressionLists();

// invoke the operation and iterate over the result
await foreach (EmailSuppressionListResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EmailSuppressionListData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
