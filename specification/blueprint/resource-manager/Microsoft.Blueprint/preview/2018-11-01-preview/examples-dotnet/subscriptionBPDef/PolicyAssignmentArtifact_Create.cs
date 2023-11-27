using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;
using Azure.ResourceManager.Blueprint.Models;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/PolicyAssignmentArtifact_Create.json
// this example is just showing the usage of "Artifacts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this BlueprintArtifactResource
BlueprintArtifactCollection collection = blueprint.GetBlueprintArtifacts();

// invoke the operation
string artifactName = "costCenterPolicy";
ArtifactData data = new PolicyAssignmentArtifact("/providers/Microsoft.Authorization/policyDefinitions/1e30110a-5ceb-460c-a204-c1c3969c6d62", new Dictionary<string, ParameterValue>()
{
    ["tagName"] = new ParameterValue()
    {
        Value = BinaryData.FromString("\"costCenter\""),
    },
    ["tagValue"] = new ParameterValue()
    {
        Value = BinaryData.FromString("\"[parameter('costCenter')]\""),
    },
})
{
    DisplayName = "force costCenter tag on all resources",
};
ArmOperation<BlueprintArtifactResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, artifactName, data);
BlueprintArtifactResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArtifactData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
