using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiRelease.json
// this example is just showing the usage of "ApiRelease_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiReleaseResource created on azure
// for more information of creating ApiReleaseResource, please refer to the document of ApiReleaseResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "a1";
string releaseId = "5a7cb545298324c53224a799";
ResourceIdentifier apiReleaseResourceId = ApiReleaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, releaseId);
ApiReleaseResource apiRelease = client.GetApiReleaseResource(apiReleaseResourceId);

// invoke the operation
bool result = await apiRelease.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
