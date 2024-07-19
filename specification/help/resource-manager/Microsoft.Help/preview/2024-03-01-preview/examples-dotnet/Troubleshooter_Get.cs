using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_Get.json
// this example is just showing the usage of "Troubleshooters_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this SelfHelpTroubleshooterResource
string scope = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
SelfHelpTroubleshooterCollection collection = client.GetSelfHelpTroubleshooters(scopeId);

// invoke the operation
string troubleshooterName = "abf168ed-1b54-454a-86f6-e4b62253d3b1";
NullableResponse<SelfHelpTroubleshooterResource> response = await collection.GetIfExistsAsync(troubleshooterName);
SelfHelpTroubleshooterResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SelfHelpTroubleshooterData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
