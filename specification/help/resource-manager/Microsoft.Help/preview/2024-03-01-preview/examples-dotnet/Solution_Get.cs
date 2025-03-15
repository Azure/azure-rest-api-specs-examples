using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SelfHelp.Models;
using Azure.ResourceManager.SelfHelp;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Solution_Get.json
// this example is just showing the usage of "Solution_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SelfHelpSolutionResource created on azure
// for more information of creating SelfHelpSolutionResource, please refer to the document of SelfHelpSolutionResource
string scope = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
string solutionResourceName = "SolutionResource1";
ResourceIdentifier selfHelpSolutionResourceId = SelfHelpSolutionResource.CreateResourceIdentifier(scope, solutionResourceName);
SelfHelpSolutionResource selfHelpSolution = client.GetSelfHelpSolutionResource(selfHelpSolutionResourceId);

// invoke the operation
SelfHelpSolutionResource result = await selfHelpSolution.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SelfHelpSolutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
