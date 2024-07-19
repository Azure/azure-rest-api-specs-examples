using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiSchema.json
// this example is just showing the usage of "ApiSchema_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiSchemaResource created on azure
// for more information of creating ApiSchemaResource, please refer to the document of ApiSchemaResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "59d5b28d1f7fab116c282650";
string schemaId = "59d5b28e1f7fab116402044e";
ResourceIdentifier apiSchemaResourceId = ApiSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, schemaId);
ApiSchemaResource apiSchema = client.GetApiSchemaResource(apiSchemaResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await apiSchema.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine($"Succeeded");
