using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiCenter;

// Generated from example definition: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/MetadataSchemas_Head.json
// this example is just showing the usage of "MetadataSchemas_Head" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiCenterMetadataSchemaResource created on azure
// for more information of creating ApiCenterMetadataSchemaResource, please refer to the document of ApiCenterMetadataSchemaResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso-resources";
string serviceName = "contoso";
string metadataSchemaName = "author";
ResourceIdentifier apiCenterMetadataSchemaResourceId = ApiCenterMetadataSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, metadataSchemaName);
ApiCenterMetadataSchemaResource apiCenterMetadataSchema = client.GetApiCenterMetadataSchemaResource(apiCenterMetadataSchemaResourceId);

// invoke the operation
bool result = await apiCenterMetadataSchema.HeadAsync();

Console.WriteLine($"Succeeded: {result}");
