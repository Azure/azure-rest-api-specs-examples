using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Blueprint.Models;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/SealedBlueprint_List.json
// this example is just showing the usage of "PublishedBlueprints_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlueprintResource created on azure
// for more information of creating BlueprintResource, please refer to the document of BlueprintResource
string resourceScope = "subscriptions/00000000-0000-0000-0000-000000000000";
string blueprintName = "simpleBlueprint";
ResourceIdentifier blueprintResourceId = BlueprintResource.CreateResourceIdentifier(resourceScope, blueprintName);
BlueprintResource blueprint = client.GetBlueprintResource(blueprintResourceId);

// invoke the operation and iterate over the result
await foreach (PublishedBlueprintResource item in blueprint.GetPublishedBlueprintsAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PublishedBlueprintData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
