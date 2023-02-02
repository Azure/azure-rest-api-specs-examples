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

// Generated from example definition: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_Update.json
// this example is just showing the usage of "LoadTests_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LoadTestingResource created on azure
// for more information of creating LoadTestingResource, please refer to the document of LoadTestingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "dummyrg";
string loadTestName = "myLoadTest";
ResourceIdentifier loadTestingResourceId = LoadTestingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadTestName);
LoadTestingResource loadTestingResource = client.GetLoadTestingResource(loadTestingResourceId);

// invoke the operation
LoadTestingResourcePatch patch = new LoadTestingResourcePatch()
{
    Tags =
    {
    ["Division"] = "LT",
    ["Team"] = "Dev Exp",
    },
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
            IdentityType = LoadTestingCmkIdentityType.SystemAssigned,
            ResourceId = null,
        },
        KeyUri = new Uri("https://dummy.vault.azure.net/keys/dummykey1"),
    },
};
ArmOperation<LoadTestingResource> lro = await loadTestingResource.UpdateAsync(WaitUntil.Completed, patch);
LoadTestingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LoadTestingResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
