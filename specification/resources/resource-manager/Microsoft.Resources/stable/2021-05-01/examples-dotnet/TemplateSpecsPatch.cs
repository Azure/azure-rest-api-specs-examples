using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecsPatch.json
// this example is just showing the usage of "TemplateSpecs_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TemplateSpecResource created on azure
// for more information of creating TemplateSpecResource, please refer to the document of TemplateSpecResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "templateSpecRG";
string templateSpecName = "simpleTemplateSpec";
ResourceIdentifier templateSpecResourceId = TemplateSpecResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, templateSpecName);
TemplateSpecResource templateSpec = client.GetTemplateSpecResource(templateSpecResourceId);

// invoke the operation
TemplateSpecPatch patch = new TemplateSpecPatch()
{
    Tags =
    {
    ["myTag"] = "My Value",
    },
};
TemplateSpecResource result = await templateSpec.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TemplateSpecData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
