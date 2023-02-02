using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LoadTesting;
using Azure.ResourceManager.LoadTesting.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_CreateOrUpdate.json
// this example is just showing the usage of "LoadTests_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "dummyrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this LoadTestingResource
LoadTestingResourceCollection collection = resourceGroupResource.GetLoadTestingResources();

// invoke the operation
string loadTestName = "myLoadTest";
LoadTestingResourceData data = new LoadTestingResourceData(new AzureLocation("westus"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned,UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity(),
        },
    },
    Description = "This is new load test resource",
    Encryption = new LoadTestingCmkEncryptionProperties()
    {
        Identity = new LoadTestingCmkIdentity()
        {
            IdentityType = LoadTestingCmkIdentityType.UserAssigned,
            ResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
        },
        KeyUri = new Uri("https://dummy.vault.azure.net/keys/dummykey1"),
    },
    Tags =
    {
    ["Team"] = "Dev Exp",
    },
};
ArmOperation<LoadTestingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, loadTestName, data);
LoadTestingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LoadTestingResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
