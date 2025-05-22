using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteWorkspaceApiSchema.json
// this example is just showing the usage of "WorkspaceApiSchema_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceApiSchemaResource created on azure
// for more information of creating ServiceWorkspaceApiSchemaResource, please refer to the document of ServiceWorkspaceApiSchemaResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string apiId = "59d5b28d1f7fab116c282650";
string schemaId = "59d5b28e1f7fab116402044e";
ResourceIdentifier serviceWorkspaceApiSchemaResourceId = ServiceWorkspaceApiSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, apiId, schemaId);
ServiceWorkspaceApiSchemaResource serviceWorkspaceApiSchema = client.GetServiceWorkspaceApiSchemaResource(serviceWorkspaceApiSchemaResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await serviceWorkspaceApiSchema.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine("Succeeded");
