using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: 2024-09-01-preview/Delete_Schema.json
// this example is just showing the usage of "Schema_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceRegistrySchemaResource created on azure
// for more information of creating DeviceRegistrySchemaResource, please refer to the document of DeviceRegistrySchemaResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string schemaRegistryName = "my-schema-registry";
string schemaName = "my-schema";
ResourceIdentifier deviceRegistrySchemaResourceId = DeviceRegistrySchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schemaRegistryName, schemaName);
DeviceRegistrySchemaResource deviceRegistrySchema = client.GetDeviceRegistrySchemaResource(deviceRegistrySchemaResourceId);

// invoke the operation
await deviceRegistrySchema.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
