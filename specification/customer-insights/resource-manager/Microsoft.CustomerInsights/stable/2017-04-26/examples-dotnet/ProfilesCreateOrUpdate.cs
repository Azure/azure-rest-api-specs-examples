using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CustomerInsights.Models;
using Azure.ResourceManager.CustomerInsights;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfilesCreateOrUpdate.json
// this example is just showing the usage of "Profiles_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProfileResourceFormatResource created on azure
// for more information of creating ProfileResourceFormatResource, please refer to the document of ProfileResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string profileName = "TestProfileType396";
ResourceIdentifier profileResourceFormatResourceId = ProfileResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, profileName);
ProfileResourceFormatResource profileResourceFormat = client.GetProfileResourceFormatResource(profileResourceFormatResourceId);

// invoke the operation
ProfileResourceFormatData data = new ProfileResourceFormatData
{
    SmallImage = "\\\\Images\\\\smallImage",
    MediumImage = "\\\\Images\\\\MediumImage",
    LargeImage = "\\\\Images\\\\LargeImage",
    ApiEntitySetName = "TestProfileType396",
    Fields = {new PropertyDefinition("Id", "Edm.String")
    {
    IsArray = false,
    IsRequired = true,
    }, new PropertyDefinition("ProfileId", "Edm.String")
    {
    IsArray = false,
    IsRequired = true,
    }, new PropertyDefinition("LastName", "Edm.String")
    {
    IsArray = false,
    IsRequired = true,
    }, new PropertyDefinition("TestProfileType396", "Edm.String")
    {
    IsArray = false,
    IsRequired = true,
    }, new PropertyDefinition("SavingAccountBalance", "Edm.Int32")
    {
    IsArray = false,
    IsRequired = true,
    }},
    SchemaItemTypeLink = "SchemaItemTypeLink",
    StrongIds = { new StrongId(new string[] { "Id", "SavingAccountBalance" }, "Id"), new StrongId(new string[] { "ProfileId", "LastName" }, "ProfileId") },
};
ArmOperation<ProfileResourceFormatResource> lro = await profileResourceFormat.UpdateAsync(WaitUntil.Completed, data);
ProfileResourceFormatResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProfileResourceFormatData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
