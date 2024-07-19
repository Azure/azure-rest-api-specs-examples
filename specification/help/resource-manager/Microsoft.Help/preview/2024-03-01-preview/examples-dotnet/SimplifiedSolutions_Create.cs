using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/SimplifiedSolutions_Create.json
// this example is just showing the usage of "SimplifiedSolutions_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SelfHelpSimplifiedSolutionResource created on azure
// for more information of creating SelfHelpSimplifiedSolutionResource, please refer to the document of SelfHelpSimplifiedSolutionResource
string scope = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
string simplifiedSolutionsResourceName = "simplifiedSolutionsResourceName1";
ResourceIdentifier selfHelpSimplifiedSolutionResourceId = SelfHelpSimplifiedSolutionResource.CreateResourceIdentifier(scope, simplifiedSolutionsResourceName);
SelfHelpSimplifiedSolutionResource selfHelpSimplifiedSolution = client.GetSelfHelpSimplifiedSolutionResource(selfHelpSimplifiedSolutionResourceId);

// invoke the operation
SelfHelpSimplifiedSolutionData data = new SelfHelpSimplifiedSolutionData()
{
    SolutionId = "sampleSolutionId",
    Parameters =
    {
    ["resourceUri"] = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
    },
};
ArmOperation<SelfHelpSimplifiedSolutionResource> lro = await selfHelpSimplifiedSolution.UpdateAsync(WaitUntil.Completed, data);
SelfHelpSimplifiedSolutionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SelfHelpSimplifiedSolutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
