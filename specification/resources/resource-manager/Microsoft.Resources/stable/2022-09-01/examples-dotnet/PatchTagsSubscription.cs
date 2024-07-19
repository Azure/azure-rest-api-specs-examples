using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/PatchTagsSubscription.json
// this example is just showing the usage of "Tags_UpdateAtScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TagResource created on azure
// for more information of creating TagResource, please refer to the document of TagResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000";
ResourceIdentifier tagResourceId = TagResource.CreateResourceIdentifier(scope);
TagResource tagResource = client.GetTagResource(tagResourceId);

// invoke the operation
TagResourcePatch patch = new TagResourcePatch()
{
    PatchMode = TagPatchMode.Replace,
    TagValues =
    {
    ["tagKey1"] = "tag-value-1",
    ["tagKey2"] = "tag-value-2",
    },
};
ArmOperation<TagResource> lro = await tagResource.UpdateAsync(WaitUntil.Completed, patch);
TagResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TagResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
