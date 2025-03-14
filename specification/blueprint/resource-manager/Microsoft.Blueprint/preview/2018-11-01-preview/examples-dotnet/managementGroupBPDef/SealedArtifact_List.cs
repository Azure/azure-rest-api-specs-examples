using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/SealedArtifact_List.json
// this example is just showing the usage of "PublishedArtifacts_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublishedBlueprintResource created on azure
// for more information of creating PublishedBlueprintResource, please refer to the document of PublishedBlueprintResource
string resourceScope = "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup";
string blueprintName = "simpleBlueprint";
string versionId = "V2";
ResourceIdentifier publishedBlueprintResourceId = PublishedBlueprintResource.CreateResourceIdentifier(resourceScope, blueprintName, versionId);
PublishedBlueprintResource publishedBlueprint = client.GetPublishedBlueprintResource(publishedBlueprintResourceId);

// get the collection of this BlueprintVersionArtifactResource
BlueprintVersionArtifactCollection collection = publishedBlueprint.GetBlueprintVersionArtifacts();

// invoke the operation and iterate over the result
await foreach (BlueprintVersionArtifactResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ArtifactData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
