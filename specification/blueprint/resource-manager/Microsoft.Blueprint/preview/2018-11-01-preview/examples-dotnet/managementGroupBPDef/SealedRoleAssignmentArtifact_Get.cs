using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/SealedRoleAssignmentArtifact_Get.json
// this example is just showing the usage of "PublishedArtifacts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlueprintVersionArtifactResource created on azure
// for more information of creating BlueprintVersionArtifactResource, please refer to the document of BlueprintVersionArtifactResource
string resourceScope = "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup";
string blueprintName = "simpleBlueprint";
string versionId = "V2";
string artifactName = "ownerAssignment";
ResourceIdentifier blueprintVersionArtifactResourceId = BlueprintVersionArtifactResource.CreateResourceIdentifier(resourceScope, blueprintName, versionId, artifactName);
BlueprintVersionArtifactResource blueprintVersionArtifact = client.GetBlueprintVersionArtifactResource(blueprintVersionArtifactResourceId);

// invoke the operation
BlueprintVersionArtifactResource result = await blueprintVersionArtifact.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArtifactData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
