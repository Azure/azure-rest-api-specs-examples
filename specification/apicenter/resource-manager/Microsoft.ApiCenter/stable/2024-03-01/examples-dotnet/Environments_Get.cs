using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiCenter;

// Generated from example definition: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Environments_Get.json
// this example is just showing the usage of "Environments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiCenterEnvironmentResource created on azure
// for more information of creating ApiCenterEnvironmentResource, please refer to the document of ApiCenterEnvironmentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso-resources";
string serviceName = "contoso";
string workspaceName = "default";
string environmentName = "public";
ResourceIdentifier apiCenterEnvironmentResourceId = ApiCenterEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceName, environmentName);
ApiCenterEnvironmentResource apiCenterEnvironment = client.GetApiCenterEnvironmentResource(apiCenterEnvironmentResourceId);

// invoke the operation
ApiCenterEnvironmentResource result = await apiCenterEnvironment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiCenterEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
