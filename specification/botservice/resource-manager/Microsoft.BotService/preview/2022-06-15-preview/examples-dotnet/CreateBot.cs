using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/preview/2022-06-15-preview/examples/CreateBot.json
// this example is just showing the usage of "Bots_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subscription-id";
string resourceGroupName = "OneResourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this BotResource
BotCollection collection = resourceGroupResource.GetBots();

// invoke the operation
string resourceName = "samplebotname";
BotData data = new BotData(new AzureLocation("West US"))
{
    Properties = new BotProperties("The Name of the bot", "http://mybot.coffee", "exampleappid")
    {
        Description = "The description of the bot",
        IconUri = new Uri("http://myicon"),
        MsaAppType = MsaAppType.UserAssignedMSI,
        MsaAppTenantId = "exampleapptenantid",
        MsaAppMSIResourceId = "/subscriptions/foo/resourcegroups/bar/providers/microsoft.managedidentity/userassignedidentities/sampleId",
        DeveloperAppInsightKey = "appinsightskey",
        DeveloperAppInsightsApiKey = "appinsightsapikey",
        DeveloperAppInsightsApplicationId = "appinsightsappid",
        LuisAppIds =
        {
        "luisappid1","luisappid2"
        },
        LuisKey = "luiskey",
        IsCmekEnabled = true,
        CmekKeyVaultUri = new Uri("https://myCmekKey"),
        PublicNetworkAccess = PublicNetworkAccess.Enabled,
        DisableLocalAuth = true,
        SchemaTransformationVersion = "1.0",
    },
    Sku = new BotServiceSku(BotServiceSkuName.S1),
    Kind = BotServiceKind.Sdk,
    ETag = new ETag("etag1"),
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
ArmOperation<BotResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
BotResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BotData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
