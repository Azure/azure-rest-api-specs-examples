using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/ProjectEnvironmentTypes_List.json
// this example is just showing the usage of "ProjectEnvironmentTypes_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProjectResource created on azure
// for more information of creating ProjectResource, please refer to the document of ProjectResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string projectName = "ContosoProj";
ResourceIdentifier projectResourceId = ProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName);
ProjectResource project = client.GetProjectResource(projectResourceId);

// get the collection of this ProjectEnvironmentTypeResource
ProjectEnvironmentTypeCollection collection = project.GetProjectEnvironmentTypes();

// invoke the operation and iterate over the result
await foreach (ProjectEnvironmentTypeResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ProjectEnvironmentTypeData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
