using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SelfHelp;
using Azure.ResourceManager.SelfHelp.Models;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Troubleshooter_Get.json
// this example is just showing the usage of "Troubleshooters_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TroubleshooterResource created on azure
// for more information of creating TroubleshooterResource, please refer to the document of TroubleshooterResource
string scope = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
string troubleshooterName = "abf168ed-1b54-454a-86f6-e4b62253d3b1";
ResourceIdentifier troubleshooterResourceId = TroubleshooterResource.CreateResourceIdentifier(scope, troubleshooterName);
TroubleshooterResource troubleshooterResource = client.GetTroubleshooterResource(troubleshooterResourceId);

// invoke the operation
TroubleshooterResource result = await troubleshooterResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TroubleshooterResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
