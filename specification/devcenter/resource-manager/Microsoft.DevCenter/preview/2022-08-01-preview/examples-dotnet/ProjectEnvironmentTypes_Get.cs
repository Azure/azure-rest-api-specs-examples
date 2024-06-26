using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/ProjectEnvironmentTypes_Get.json
// this example is just showing the usage of "ProjectEnvironmentTypes_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string environmentTypeName = "{environmentTypeName}";
bool result = await collection.ExistsAsync(environmentTypeName);

Console.WriteLine($"Succeeded: {result}");
