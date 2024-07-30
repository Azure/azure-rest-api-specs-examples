using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/AppServiceEnvironments_GetMultiRolePool.json
// this example is just showing the usage of "AppServiceEnvironments_GetMultiRolePool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HostingEnvironmentMultiRolePoolResource created on azure
// for more information of creating HostingEnvironmentMultiRolePoolResource, please refer to the document of HostingEnvironmentMultiRolePoolResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-rg";
string name = "test-ase";
ResourceIdentifier hostingEnvironmentMultiRolePoolResourceId = HostingEnvironmentMultiRolePoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
HostingEnvironmentMultiRolePoolResource hostingEnvironmentMultiRolePool = client.GetHostingEnvironmentMultiRolePoolResource(hostingEnvironmentMultiRolePoolResourceId);

// invoke the operation
HostingEnvironmentMultiRolePoolResource result = await hostingEnvironmentMultiRolePool.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppServiceWorkerPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
