using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerTopics_CreateOrUpdate.json
// this example is just showing the usage of "PartnerTopics_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PartnerTopicResource
PartnerTopicCollection collection = resourceGroupResource.GetPartnerTopics();

// invoke the operation
string partnerTopicName = "examplePartnerTopicName1";
PartnerTopicData data = new PartnerTopicData(new AzureLocation("westus2"))
{
    PartnerRegistrationImmutableId = Guid.Parse("6f541064-031d-4cc8-9ec3-a3b4fc0f7185"),
    Source = "ContosoCorp.Accounts.User1",
    ExpireOnIfNotActivated = DateTimeOffset.Parse("2022-03-23T23:06:13.109Z"),
    PartnerTopicFriendlyDescription = "Example description",
    MessageForActivation = "Example message for activation",
};
ArmOperation<PartnerTopicResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, partnerTopicName, data);
PartnerTopicResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PartnerTopicData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
