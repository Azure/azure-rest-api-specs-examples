using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/Instance_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "InstanceResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IotOperationsInstanceResource
IotOperationsInstanceCollection collection = resourceGroupResource.GetIotOperationsInstances();

// invoke the operation
string instanceName = "aio-instance";
IotOperationsInstanceData data = new IotOperationsInstanceData(new AzureLocation("eastus2"), new IotOperationsExtendedLocation("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123", IotOperationsExtendedLocationType.CustomLocation))
{
    Properties = new IotOperationsInstanceProperties(new SchemaRegistryRef(new ResourceIdentifier("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123")))
    {
        Description = "kpqtgocs",
    },
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities = { },
    },
    Tags = { },
};
ArmOperation<IotOperationsInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, instanceName, data);
IotOperationsInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
