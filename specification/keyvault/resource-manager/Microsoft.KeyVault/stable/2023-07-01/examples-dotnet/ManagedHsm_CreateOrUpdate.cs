using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.KeyVault.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.KeyVault;

// Generated from example definition: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/ManagedHsm_CreateOrUpdate.json
// this example is just showing the usage of "ManagedHsms_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "hsm-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ManagedHsmResource
ManagedHsmCollection collection = resourceGroupResource.GetManagedHsms();

// invoke the operation
string name = "hsm1";
ManagedHsmData data = new ManagedHsmData(new AzureLocation("westus"))
{
    Properties = new ManagedHsmProperties
    {
        TenantId = Guid.Parse("00000000-0000-0000-0000-000000000000"),
        InitialAdminObjectIds = { "00000000-0000-0000-0000-000000000000" },
        EnableSoftDelete = true,
        SoftDeleteRetentionInDays = 90,
        EnablePurgeProtection = false,
    },
    Sku = new ManagedHsmSku(ManagedHsmSkuFamily.B, ManagedHsmSkuName.StandardB1),
    Tags =
    {
    ["Dept"] = "hsm",
    ["Environment"] = "dogfood"
    },
};
ArmOperation<ManagedHsmResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
ManagedHsmResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedHsmData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
