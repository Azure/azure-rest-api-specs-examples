using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maintenance;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ApplyUpdates_Get.json
// this example is just showing the usage of "ApplyUpdates_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MaintenanceApplyUpdateResource
MaintenanceApplyUpdateCollection collection = resourceGroupResource.GetMaintenanceApplyUpdates();

// invoke the operation
string providerName = "Microsoft.Compute";
string resourceType = "virtualMachineScaleSets";
string resourceName = "smdtest1";
string applyUpdateName = "e9b9685d-78e4-44c4-a81c-64a14f9b87b6";
NullableResponse<MaintenanceApplyUpdateResource> response = await collection.GetIfExistsAsync(providerName, resourceType, resourceName, applyUpdateName);
MaintenanceApplyUpdateResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MaintenanceApplyUpdateData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
