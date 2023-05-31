using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;
using Azure.ResourceManager.Blueprint.Models;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/ResourceGroupWithTags.json
// this example is just showing the usage of "Blueprints_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this BlueprintResource
string resourceScope = "providers/Microsoft.Management/managementGroups/{ManagementGroupId}";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceScope));
BlueprintCollection collection = client.GetBlueprints(scopeId);

// invoke the operation
string blueprintName = "simpleBlueprint";
BlueprintData data = new BlueprintData()
{
    Description = "An example blueprint containing an RG with two tags.",
    TargetScope = BlueprintTargetScope.Subscription,
    ResourceGroups =
    {
    ["myRGName"] = new ResourceGroupDefinition()
    {
    Name = "myRGName",
    Location = new AzureLocation("westus"),
    Tags =
    {
    ["costcenter"] = "123456",
    ["nameOnlyTag"] = "",
    },
    DisplayName = "My Resource Group",
    },
    },
};
ArmOperation<BlueprintResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, blueprintName, data);
BlueprintResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BlueprintData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
