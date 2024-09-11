using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementHeadProductApi.json
// this example is just showing the usage of "ProductApi_CheckEntityExists" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementProductResource created on azure
// for more information of creating ApiManagementProductResource, please refer to the document of ApiManagementProductResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string productId = "5931a75ae4bbd512a88c680b";
ResourceIdentifier apiManagementProductResourceId = ApiManagementProductResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, productId);
ApiManagementProductResource apiManagementProduct = client.GetApiManagementProductResource(apiManagementProductResourceId);

// invoke the operation
string apiId = "59306a29e4bbd510dc24e5f9";
bool result = await apiManagementProduct.CheckProductApiEntityExistsAsync(apiId);

Console.WriteLine($"Succeeded: {result}");
