using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SelfHelp.Models;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_Create.json
// this example is just showing the usage of "Troubleshooters_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SelfHelpTroubleshooterResource created on azure
// for more information of creating SelfHelpTroubleshooterResource, please refer to the document of SelfHelpTroubleshooterResource
string scope = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
string troubleshooterName = "abf168ed-1b54-454a-86f6-e4b62253d3b1";
ResourceIdentifier selfHelpTroubleshooterResourceId = SelfHelpTroubleshooterResource.CreateResourceIdentifier(scope, troubleshooterName);
SelfHelpTroubleshooterResource selfHelpTroubleshooter = client.GetSelfHelpTroubleshooterResource(selfHelpTroubleshooterResourceId);

// invoke the operation
SelfHelpTroubleshooterData data = new SelfHelpTroubleshooterData
{
    SolutionId = "SampleTroubleshooterSolutionId",
    Parameters =
    {
    ["ResourceURI"] = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp"
    },
};
ArmOperation<SelfHelpTroubleshooterResource> lro = await selfHelpTroubleshooter.UpdateAsync(WaitUntil.Completed, data);
SelfHelpTroubleshooterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SelfHelpTroubleshooterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
