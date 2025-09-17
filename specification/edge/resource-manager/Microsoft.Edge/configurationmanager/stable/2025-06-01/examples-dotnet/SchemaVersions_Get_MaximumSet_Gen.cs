using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/SchemaVersions_Get_MaximumSet_Gen.json
// this example is just showing the usage of "SchemaVersion_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeSchemaResource created on azure
// for more information of creating EdgeSchemaResource, please refer to the document of EdgeSchemaResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string schemaName = "testname";
ResourceIdentifier edgeSchemaResourceId = EdgeSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schemaName);
EdgeSchemaResource edgeSchema = client.GetEdgeSchemaResource(edgeSchemaResourceId);

// get the collection of this EdgeSchemaVersionResource
EdgeSchemaVersionCollection collection = edgeSchema.GetEdgeSchemaVersions();

// invoke the operation
string schemaVersionName = "1.0.0";
NullableResponse<EdgeSchemaVersionResource> response = await collection.GetIfExistsAsync(schemaVersionName);
EdgeSchemaVersionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EdgeSchemaVersionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
