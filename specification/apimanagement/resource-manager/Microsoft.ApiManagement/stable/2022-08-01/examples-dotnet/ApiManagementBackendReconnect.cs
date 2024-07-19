using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackendReconnect.json
// this example is just showing the usage of "Backend_Reconnect" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementBackendResource created on azure
// for more information of creating ApiManagementBackendResource, please refer to the document of ApiManagementBackendResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string backendId = "proxybackend";
ResourceIdentifier apiManagementBackendResourceId = ApiManagementBackendResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, backendId);
ApiManagementBackendResource apiManagementBackend = client.GetApiManagementBackendResource(apiManagementBackendResourceId);

// invoke the operation
BackendReconnectContract backendReconnectContract = new BackendReconnectContract()
{
    After = XmlConvert.ToTimeSpan("PT3S"),
};
await apiManagementBackend.ReconnectAsync(backendReconnectContract: backendReconnectContract);

Console.WriteLine($"Succeeded");
