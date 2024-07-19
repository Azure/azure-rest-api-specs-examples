using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiRelease.json
// this example is just showing the usage of "ApiRelease_Get" operation, for the dependent resources, they will have to be created separately.

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
ApiReleaseResource result = await apiRelease.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiReleaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
