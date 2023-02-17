using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;
using Azure.ResourceManager.Blueprint.Models;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/RoleAssignmentArtifact_Delete.json
// this example is just showing the usage of "Artifacts_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlueprintArtifactResource created on azure
// for more information of creating BlueprintArtifactResource, please refer to the document of BlueprintArtifactResource
string resourceScope = "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup";
string blueprintName = "simpleBlueprint";
string artifactName = "ownerAssignment";
ResourceIdentifier blueprintArtifactResourceId = BlueprintArtifactResource.CreateResourceIdentifier(resourceScope, blueprintName, artifactName);
BlueprintArtifactResource blueprintArtifact = client.GetBlueprintArtifactResource(blueprintArtifactResourceId);

// invoke the operation
ArmOperation<BlueprintArtifactResource> lro = await blueprintArtifact.DeleteAsync(WaitUntil.Completed);
BlueprintArtifactResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArtifactData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
