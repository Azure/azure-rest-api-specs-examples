using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/DeleteSharedCommitmentPlanAssociation.json
// this example is just showing the usage of "CommitmentPlans_DeleteAssociation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CommitmentPlanAccountAssociationResource created on azure
// for more information of creating CommitmentPlanAccountAssociationResource, please refer to the document of CommitmentPlanAccountAssociationResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
string resourceGroupName = "resourceGroupName";
string commitmentPlanName = "commitmentPlanName";
string commitmentPlanAssociationName = "commitmentPlanAssociationName";
ResourceIdentifier commitmentPlanAccountAssociationResourceId = CommitmentPlanAccountAssociationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, commitmentPlanName, commitmentPlanAssociationName);
CommitmentPlanAccountAssociationResource commitmentPlanAccountAssociation = client.GetCommitmentPlanAccountAssociationResource(commitmentPlanAccountAssociationResourceId);

// invoke the operation
await commitmentPlanAccountAssociation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
