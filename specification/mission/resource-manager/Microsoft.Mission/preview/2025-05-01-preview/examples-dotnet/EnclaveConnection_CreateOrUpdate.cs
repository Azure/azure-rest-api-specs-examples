using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/EnclaveConnection_CreateOrUpdate.json
// this example is just showing the usage of "EnclaveConnectionResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "73CEECEF-2C30-488E-946F-D20F414D99BA";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VirtualEnclaveConnectionResource
VirtualEnclaveConnectionCollection collection = resourceGroupResource.GetVirtualEnclaveConnections();

// invoke the operation
string enclaveConnectionName = "TestMyEnclaveConnection";
VirtualEnclaveConnectionData data = new VirtualEnclaveConnectionData(new AzureLocation("West US"))
{
    Properties = new VirtualEnclaveConnectionProperties(new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/testrg/providers/Microsoft.Mission/communities/TestMyCommunity"), new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/microsoft.mission/virtualenclaves/TestMyEnclave"), new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/virtualenclaves/TestMyEnclave/enclaveendpoints/TestMyEnclaveEndpoint"))
    {
        SourceCidr = "10.0.0.0/24",
    },
    Tags =
    {
    ["sampletag"] = "samplevalue"
    },
};
ArmOperation<VirtualEnclaveConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, enclaveConnectionName, data);
VirtualEnclaveConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
