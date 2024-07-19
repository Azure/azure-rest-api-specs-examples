using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/DeleteCommitmentPlan.json
// this example is just showing the usage of "CommitmentPlans_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CommitmentPlanResource created on azure
// for more information of creating CommitmentPlanResource, please refer to the document of CommitmentPlanResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string commitmentPlanName = "commitmentPlanName";
ResourceIdentifier commitmentPlanResourceId = CommitmentPlanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, commitmentPlanName);
CommitmentPlanResource commitmentPlan = client.GetCommitmentPlanResource(commitmentPlanResourceId);

// invoke the operation
await commitmentPlan.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
