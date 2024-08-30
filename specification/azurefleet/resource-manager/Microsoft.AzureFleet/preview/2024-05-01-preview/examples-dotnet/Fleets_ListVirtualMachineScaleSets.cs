using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeFleet.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ComputeFleet;

// Generated from example definition: specification/azurefleet/resource-manager/Microsoft.AzureFleet/preview/2024-05-01-preview/examples/Fleets_ListVirtualMachineScaleSets.json
// this example is just showing the usage of "Fleets_ListVirtualMachineScaleSets" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ComputeFleetResource created on azure
// for more information of creating ComputeFleetResource, please refer to the document of ComputeFleetResource
string subscriptionId = "1DC2F28C-A625-4B0E-9748-9885A3C9E9EB";
string resourceGroupName = "rgazurefleet";
string name = "myFleet";
ResourceIdentifier computeFleetResourceId = ComputeFleetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
ComputeFleetResource computeFleet = client.GetComputeFleetResource(computeFleetResourceId);

// invoke the operation and iterate over the result
await foreach (ComputeFleetVmss item in computeFleet.GetVirtualMachineScaleSetsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
