using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiSchema.json
// this example is just showing the usage of "ApiSchema_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiSchemaResource created on azure
// for more information of creating ApiSchemaResource, please refer to the document of ApiSchemaResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "57d1f7558aa04f15146d9d8a";
string schemaId = "ec12520d-9d48-4e7b-8f39-698ca2ac63f1";
ResourceIdentifier apiSchemaResourceId = ApiSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, schemaId);
ApiSchemaResource apiSchema = client.GetApiSchemaResource(apiSchemaResourceId);

// invoke the operation
bool result = await apiSchema.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
