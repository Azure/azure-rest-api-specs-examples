using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/HealthModels_Update.json
// this example is just showing the usage of "HealthModel_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthModelResource created on azure
// for more information of creating HealthModelResource, please refer to the document of HealthModelResource
string subscriptionId = "4980D7D5-4E07-47AD-AD34-E76C6BC9F061";
string resourceGroupName = "rgopenapi";
string healthModelName = "model1";
ResourceIdentifier healthModelResourceId = HealthModelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, healthModelName);
HealthModelResource healthModel = client.GetHealthModelResource(healthModelResourceId);

// invoke the operation
HealthModelPatch patch = new HealthModelPatch
{
    Identity = new ManagedServiceIdentity("SystemAssigned, UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ua1")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["key21"] = "menfkmseplchh"
    },
};
ArmOperation<HealthModelResource> lro = await healthModel.UpdateAsync(WaitUntil.Completed, patch);
HealthModelResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthModelData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
