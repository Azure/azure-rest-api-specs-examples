using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiSchema.json
// this example is just showing the usage of "ApiSchema_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiSchemaResource created on azure
// for more information of creating ApiSchemaResource, please refer to the document of ApiSchemaResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "59d6bb8f1f7fab13dc67ec9b";
string schemaId = "ec12520d-9d48-4e7b-8f39-698ca2ac63f1";
ResourceIdentifier apiSchemaResourceId = ApiSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, schemaId);
ApiSchemaResource apiSchema = client.GetApiSchemaResource(apiSchemaResourceId);

// invoke the operation
ApiSchemaResource result = await apiSchema.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiSchemaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
