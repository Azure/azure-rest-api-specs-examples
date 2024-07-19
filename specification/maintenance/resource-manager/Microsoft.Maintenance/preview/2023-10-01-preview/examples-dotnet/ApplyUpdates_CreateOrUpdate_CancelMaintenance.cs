using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Maintenance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maintenance;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ApplyUpdates_CreateOrUpdate_CancelMaintenance.json
// this example is just showing the usage of "ApplyUpdates_CreateOrUpdateOrCancel" operation, for the dependent resources, they will have to be created separately.

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
string providerName = "Microsoft.Maintenance";
string resourceType = "maintenanceConfigurations";
string resourceName = "maintenanceConfig1";
string applyUpdateName = "20230901121200";
MaintenanceApplyUpdateData data = new MaintenanceApplyUpdateData()
{
    Status = MaintenanceUpdateStatus.Cancel,
};
ArmOperation<MaintenanceApplyUpdateResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, providerName, resourceType, resourceName, applyUpdateName, data);
MaintenanceApplyUpdateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MaintenanceApplyUpdateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
