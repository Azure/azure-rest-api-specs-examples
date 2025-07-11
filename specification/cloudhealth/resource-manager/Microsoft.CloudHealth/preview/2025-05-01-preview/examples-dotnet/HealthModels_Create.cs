using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/HealthModels_Create.json
// this example is just showing the usage of "HealthModel_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "4980D7D5-4E07-47AD-AD34-E76C6BC9F061";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HealthModelResource
HealthModelCollection collection = resourceGroupResource.GetHealthModels();

// invoke the operation
string healthModelName = "model1";
HealthModelData data = new HealthModelData(new AzureLocation("eastus2"))
{
    Properties = new HealthModelProperties
    {
        Discovery = new ModelDiscoverySettings("/providers/Microsoft.Management/serviceGroups/myServiceGroup", DiscoveryRuleRecommendedSignalsBehavior.Enabled)
        {
            Identity = "SystemAssigned",
        },
    },
    Identity = new ManagedServiceIdentity("SystemAssigned, UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ua1")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["key2961"] = "hbljozzkqrpcthsjtfkyozpwyx"
    },
};
ArmOperation<HealthModelResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, healthModelName, data);
HealthModelResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthModelData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
