using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Maps;
using Azure.ResourceManager.Maps.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/CreateAccountGen2.json
// this example is just showing the usage of "Accounts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MapsAccountResource
MapsAccountCollection collection = resourceGroupResource.GetMapsAccounts();

// invoke the operation
string accountName = "myMapsAccount";
MapsAccountData data = new MapsAccountData(new AzureLocation("eastus"), new MapsSku(MapsSkuName.G2))
{
    Kind = MapsAccountKind.Gen2,
    Properties = new MapsAccountProperties()
    {
        DisableLocalAuth = true,
        CorsRulesValue =
        {
        new MapsCorsRule(new string[]
        {
        "http://www.contoso.com","http://www.fabrikam.com"
        })
        },
    },
    Tags =
    {
    ["test"] = "true",
    },
};
ArmOperation<MapsAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, data);
MapsAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MapsAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
