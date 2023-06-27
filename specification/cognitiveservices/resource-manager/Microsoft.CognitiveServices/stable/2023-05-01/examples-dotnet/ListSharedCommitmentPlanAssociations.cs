using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListSharedCommitmentPlanAssociations.json
// this example is just showing the usage of "CommitmentPlans_ListAssociations" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesCommitmentPlanResource created on azure
// for more information of creating CognitiveServicesCommitmentPlanResource, please refer to the document of CognitiveServicesCommitmentPlanResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
string resourceGroupName = "resourceGroupName";
string commitmentPlanName = "commitmentPlanName";
ResourceIdentifier cognitiveServicesCommitmentPlanResourceId = CognitiveServicesCommitmentPlanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, commitmentPlanName);
CognitiveServicesCommitmentPlanResource cognitiveServicesCommitmentPlan = client.GetCognitiveServicesCommitmentPlanResource(cognitiveServicesCommitmentPlanResourceId);

// get the collection of this CommitmentPlanAccountAssociationResource
CommitmentPlanAccountAssociationCollection collection = cognitiveServicesCommitmentPlan.GetCommitmentPlanAccountAssociations();

// invoke the operation and iterate over the result
await foreach (CommitmentPlanAccountAssociationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CommitmentPlanAccountAssociationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
