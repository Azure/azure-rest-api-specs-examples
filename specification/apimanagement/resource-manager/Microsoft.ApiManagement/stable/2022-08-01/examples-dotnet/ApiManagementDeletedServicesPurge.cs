using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeletedServicesPurge.json
// this example is just showing the usage of "DeletedServices_Purge" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementDeletedServiceResource created on azure
// for more information of creating ApiManagementDeletedServiceResource, please refer to the document of ApiManagementDeletedServiceResource
string subscriptionId = "subid";
AzureLocation location = new AzureLocation("westus");
string serviceName = "apimService3";
ResourceIdentifier apiManagementDeletedServiceResourceId = ApiManagementDeletedServiceResource.CreateResourceIdentifier(subscriptionId, location, serviceName);
ApiManagementDeletedServiceResource apiManagementDeletedService = client.GetApiManagementDeletedServiceResource(apiManagementDeletedServiceResourceId);

// invoke the operation
await apiManagementDeletedService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
