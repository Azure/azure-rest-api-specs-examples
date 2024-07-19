using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateProductApi.json
// this example is just showing the usage of "ProductApi_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementProductResource created on azure
// for more information of creating ApiManagementProductResource, please refer to the document of ApiManagementProductResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string productId = "testproduct";
ResourceIdentifier apiManagementProductResourceId = ApiManagementProductResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, productId);
ApiManagementProductResource apiManagementProduct = client.GetApiManagementProductResource(apiManagementProductResourceId);

// invoke the operation
string apiId = "echo-api";
ProductApiData result = await apiManagementProduct.CreateOrUpdateProductApiAsync(apiId);

Console.WriteLine($"Succeeded: {result}");
