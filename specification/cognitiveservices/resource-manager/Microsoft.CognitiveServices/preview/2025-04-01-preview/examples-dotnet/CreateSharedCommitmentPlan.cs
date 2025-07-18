using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/CreateSharedCommitmentPlan.json
// this example is just showing the usage of "CommitmentPlans_CreateOrUpdatePlan" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CognitiveServicesCommitmentPlanResource
CognitiveServicesCommitmentPlanCollection collection = resourceGroupResource.GetCognitiveServicesCommitmentPlans();

// invoke the operation
string commitmentPlanName = "commitmentPlanName";
CommitmentPlanData data = new CommitmentPlanData
{
    Kind = "SpeechServices",
    Sku = new CognitiveServicesSku("S0"),
    Location = new AzureLocation("West US"),
    Properties = new CommitmentPlanProperties
    {
        HostingModel = ServiceAccountHostingModel.Web,
        PlanType = "STT",
        Current = new CommitmentPeriod
        {
            Tier = "T1",
        },
        AutoRenew = true,
    },
};
ArmOperation<CognitiveServicesCommitmentPlanResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, commitmentPlanName, data);
CognitiveServicesCommitmentPlanResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CommitmentPlanData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
