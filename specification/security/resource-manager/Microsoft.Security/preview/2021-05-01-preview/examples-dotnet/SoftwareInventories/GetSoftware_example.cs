using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/SoftwareInventories/GetSoftware_example.json
// this example is just showing the usage of "SoftwareInventories_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SoftwareInventoryResource created on azure
// for more information of creating SoftwareInventoryResource, please refer to the document of SoftwareInventoryResource
string subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
string resourceGroupName = "EITAN-TESTS";
string resourceNamespace = "Microsoft.Compute";
string resourceType = "virtualMachines";
string resourceName = "Eitan-Test1";
string softwareName = "outlook_16.0.10371.20060";
ResourceIdentifier softwareInventoryResourceId = SoftwareInventoryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceNamespace, resourceType, resourceName, softwareName);
SoftwareInventoryResource softwareInventory = client.GetSoftwareInventoryResource(softwareInventoryResourceId);

// invoke the operation
SoftwareInventoryResource result = await softwareInventory.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SoftwareInventoryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
