using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Schemas_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Schema_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
EdgeSchemaPatch patch = new EdgeSchemaPatch
{
    Properties = new EdgeSchemaPatchProperties(),
    Tags =
    {
    ["key6760"] = "lknwkzihsmzbzkezkartwgsv"
    },
};
EdgeSchemaResource result = await edgeSchema.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeSchemaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
