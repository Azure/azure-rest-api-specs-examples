using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksCreateOrUpdate.json
// this example is just showing the usage of "RelationshipLinks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubResource created on azure
// for more information of creating HubResource, please refer to the document of HubResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
ResourceIdentifier hubResourceId = HubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName);
HubResource hub = client.GetHubResource(hubResourceId);

// get the collection of this RelationshipLinkResourceFormatResource
RelationshipLinkResourceFormatCollection collection = hub.GetRelationshipLinkResourceFormats();

// invoke the operation
string relationshipLinkName = "Somelink";
RelationshipLinkResourceFormatData data = new RelationshipLinkResourceFormatData()
{
    DisplayName =
    {
    ["en-us"] = "Link DisplayName",
    },
    Description =
    {
    ["en-us"] = "Link Description",
    },
    InteractionType = "testInteraction4332",
    ProfilePropertyReferences =
    {
    new ParticipantProfilePropertyReference("profile1","ProfileId")
    },
    RelatedProfilePropertyReferences =
    {
    new ParticipantProfilePropertyReference("profile1","ProfileId")
    },
    RelationshipName = "testProfile2326994",
};
ArmOperation<RelationshipLinkResourceFormatResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, relationshipLinkName, data);
RelationshipLinkResourceFormatResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RelationshipLinkResourceFormatData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
