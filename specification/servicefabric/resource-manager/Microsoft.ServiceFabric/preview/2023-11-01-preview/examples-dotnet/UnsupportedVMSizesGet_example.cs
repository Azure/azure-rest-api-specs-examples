using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabric;

// Generated from example definition: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/preview/2023-11-01-preview/examples/UnsupportedVMSizesGet_example.json
// this example is just showing the usage of "UnsupportedVmSizes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricVmSizeResource created on azure
// for more information of creating ServiceFabricVmSizeResource, please refer to the document of ServiceFabricVmSizeResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
AzureLocation location = new AzureLocation("eastus");
string vmSize = "Standard_B1ls1";
ResourceIdentifier serviceFabricVmSizeResourceId = ServiceFabricVmSizeResource.CreateResourceIdentifier(subscriptionId, location, vmSize);
ServiceFabricVmSizeResource serviceFabricVmSizeResource = client.GetServiceFabricVmSizeResource(serviceFabricVmSizeResourceId);

// invoke the operation
ServiceFabricVmSizeResource result = await serviceFabricVmSizeResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricVmSizeResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
