using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagedServiceIdentities;

// Generated from example definition: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/IdentityCreate.json
// this example is just showing the usage of "UserAssignedIdentities_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rgName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this UserAssignedIdentityResource
UserAssignedIdentityCollection collection = resourceGroupResource.GetUserAssignedIdentities();

// invoke the operation
string resourceName = "resourceName";
UserAssignedIdentityData data = new UserAssignedIdentityData(new AzureLocation("eastus"))
{
    Tags =
    {
    ["key1"] = "value1",
    ["key2"] = "value2"
    },
};
ArmOperation<UserAssignedIdentityResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
UserAssignedIdentityResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
UserAssignedIdentityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
