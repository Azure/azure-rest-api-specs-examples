using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SelfHelp.Models;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Solution_Create.json
// this example is just showing the usage of "Solution_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this SelfHelpSolutionResource
string scope = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
SelfHelpSolutionCollection collection = client.GetSelfHelpSolutions(new ResourceIdentifier(scope));

// invoke the operation
string solutionResourceName = "SolutionResourceName1";
SelfHelpSolutionData data = new SelfHelpSolutionData
{
    TriggerCriteria = {new SolutionTriggerCriterion
    {
    Name = SelfHelpName.SolutionId,
    Value = "SolutionId1",
    }},
    Parameters =
    {
    ["resourceUri"] = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp"
    },
};
ArmOperation<SelfHelpSolutionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, solutionResourceName, data);
SelfHelpSolutionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SelfHelpSolutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
