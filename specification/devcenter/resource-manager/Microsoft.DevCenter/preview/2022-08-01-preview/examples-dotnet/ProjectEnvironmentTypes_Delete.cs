using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/ProjectEnvironmentTypes_Delete.json
// this example is just showing the usage of "ProjectEnvironmentTypes_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProjectEnvironmentTypeResource created on azure
// for more information of creating ProjectEnvironmentTypeResource, please refer to the document of ProjectEnvironmentTypeResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string projectName = "ContosoProj";
string environmentTypeName = "{environmentTypeName}";
ResourceIdentifier projectEnvironmentTypeResourceId = ProjectEnvironmentTypeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, environmentTypeName);
ProjectEnvironmentTypeResource projectEnvironmentType = client.GetProjectEnvironmentTypeResource(projectEnvironmentTypeResourceId);

// invoke the operation
await projectEnvironmentType.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
