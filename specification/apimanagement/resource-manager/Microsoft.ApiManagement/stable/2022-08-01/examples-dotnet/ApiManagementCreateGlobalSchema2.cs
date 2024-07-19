using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGlobalSchema2.json
// this example is just showing the usage of "GlobalSchema_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementGlobalSchemaResource created on azure
// for more information of creating ApiManagementGlobalSchemaResource, please refer to the document of ApiManagementGlobalSchemaResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string schemaId = "schema1";
ResourceIdentifier apiManagementGlobalSchemaResourceId = ApiManagementGlobalSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, schemaId);
ApiManagementGlobalSchemaResource apiManagementGlobalSchema = client.GetApiManagementGlobalSchemaResource(apiManagementGlobalSchemaResourceId);

// invoke the operation
ApiManagementGlobalSchemaData data = new ApiManagementGlobalSchemaData()
{
    SchemaType = ApiSchemaType.Json,
    Description = "sample schema description",
    Document = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["type"] = "object",
        ["$id"] = "https://example.com/person.schema.json",
        ["$schema"] = "https://json-schema.org/draft/2020-12/schema",
        ["properties"] = new Dictionary<string, object>()
        {
            ["age"] = new Dictionary<string, object>()
            {
                ["type"] = "integer",
                ["description"] = "Age in years which must be equal to or greater than zero.",
                ["minimum"] = "0"
            },
            ["firstName"] = new Dictionary<string, object>()
            {
                ["type"] = "string",
                ["description"] = "The person's first name."
            },
            ["lastName"] = new Dictionary<string, object>()
            {
                ["type"] = "string",
                ["description"] = "The person's last name."
            }
        },
        ["title"] = "Person"
    }),
};
ArmOperation<ApiManagementGlobalSchemaResource> lro = await apiManagementGlobalSchema.UpdateAsync(WaitUntil.Completed, data);
ApiManagementGlobalSchemaResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementGlobalSchemaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
