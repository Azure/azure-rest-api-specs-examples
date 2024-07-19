using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadLogger.json
// this example is just showing the usage of "Logger_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementLoggerResource created on azure
// for more information of creating ApiManagementLoggerResource, please refer to the document of ApiManagementLoggerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string loggerId = "templateLogger";
ResourceIdentifier apiManagementLoggerResourceId = ApiManagementLoggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, loggerId);
ApiManagementLoggerResource apiManagementLogger = client.GetApiManagementLoggerResource(apiManagementLoggerResourceId);

// invoke the operation
bool result = await apiManagementLogger.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
