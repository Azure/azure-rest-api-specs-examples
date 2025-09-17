using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/ConfigTemplateVersions_ListByConfigTemplate_MaximumSet_Gen.json
// this example is just showing the usage of "ConfigTemplateVersion_ListByConfigTemplate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeConfigTemplateResource created on azure
// for more information of creating EdgeConfigTemplateResource, please refer to the document of EdgeConfigTemplateResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string configTemplateName = "testname";
ResourceIdentifier edgeConfigTemplateResourceId = EdgeConfigTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configTemplateName);
EdgeConfigTemplateResource edgeConfigTemplate = client.GetEdgeConfigTemplateResource(edgeConfigTemplateResourceId);

// get the collection of this EdgeConfigTemplateVersionResource
EdgeConfigTemplateVersionCollection collection = edgeConfigTemplate.GetEdgeConfigTemplateVersions();

// invoke the operation and iterate over the result
await foreach (EdgeConfigTemplateVersionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EdgeConfigTemplateVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
