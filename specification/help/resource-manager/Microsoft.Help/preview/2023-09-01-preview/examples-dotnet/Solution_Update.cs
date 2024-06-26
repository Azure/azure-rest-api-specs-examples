using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SelfHelp;
using Azure.ResourceManager.SelfHelp.Models;

// Generated from example definition: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Solution_Update.json
// this example is just showing the usage of "Solution_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SolutionResource created on azure
// for more information of creating SolutionResource, please refer to the document of SolutionResource
string scope = "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
string solutionResourceName = "SolutionResourceName1";
ResourceIdentifier solutionResourceId = SolutionResource.CreateResourceIdentifier(scope, solutionResourceName);
SolutionResource solutionResource = client.GetSolutionResource(solutionResourceId);

// invoke the operation
SolutionResourcePatch patch = new SolutionResourcePatch();
ArmOperation<SolutionResource> lro = await solutionResource.UpdateAsync(WaitUntil.Completed, patch);
SolutionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SolutionResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
