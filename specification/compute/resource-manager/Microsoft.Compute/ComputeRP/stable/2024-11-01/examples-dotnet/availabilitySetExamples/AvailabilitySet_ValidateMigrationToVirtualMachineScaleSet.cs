using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/availabilitySetExamples/AvailabilitySet_ValidateMigrationToVirtualMachineScaleSet.json
// this example is just showing the usage of "AvailabilitySets_ValidateMigrationToVirtualMachineScaleSet" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvailabilitySetResource created on azure
// for more information of creating AvailabilitySetResource, please refer to the document of AvailabilitySetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string availabilitySetName = "myAvailabilitySet";
ResourceIdentifier availabilitySetResourceId = AvailabilitySetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, availabilitySetName);
AvailabilitySetResource availabilitySet = client.GetAvailabilitySetResource(availabilitySetResourceId);

// invoke the operation
MigrateToVirtualMachineScaleSetInput input = new MigrateToVirtualMachineScaleSetInput(new WritableSubResource
{
    Id = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
});
await availabilitySet.ValidateMigrationToVirtualMachineScaleSetAsync(input);

Console.WriteLine("Succeeded");
