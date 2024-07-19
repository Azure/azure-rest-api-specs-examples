using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadNamedValue.json
// this example is just showing the usage of "NamedValue_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementNamedValueResource created on azure
// for more information of creating ApiManagementNamedValueResource, please refer to the document of ApiManagementNamedValueResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string namedValueId = "testarmTemplateproperties2";
ResourceIdentifier apiManagementNamedValueResourceId = ApiManagementNamedValueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, namedValueId);
ApiManagementNamedValueResource apiManagementNamedValue = client.GetApiManagementNamedValueResource(apiManagementNamedValueResourceId);

// invoke the operation
bool result = await apiManagementNamedValue.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
