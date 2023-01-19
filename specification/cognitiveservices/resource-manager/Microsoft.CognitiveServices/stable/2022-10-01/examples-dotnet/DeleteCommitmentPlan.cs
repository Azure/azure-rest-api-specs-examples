using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/DeleteCommitmentPlan.json
// this example is just showing the usage of "CommitmentPlans_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

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
