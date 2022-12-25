using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maintenance;

// Generated from example definition: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/ApplyUpdates_GetParent.json
// this example is just showing the usage of "ApplyUpdates_GetParent" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
string providerName = "Microsoft.Compute";
string resourceParentType = "virtualMachineScaleSets";
string resourceParentName = "smdtest1";
string resourceType = "virtualMachines";
string resourceName = "smdvm1";
string applyUpdateName = "e9b9685d-78e4-44c4-a81c-64a14f9b87b6";
MaintenanceApplyUpdateResource result = await resourceGroupResource.GetApplyUpdatesByParentAsync(providerName, resourceParentType, resourceParentName, resourceType, resourceName, applyUpdateName);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MaintenanceApplyUpdateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
