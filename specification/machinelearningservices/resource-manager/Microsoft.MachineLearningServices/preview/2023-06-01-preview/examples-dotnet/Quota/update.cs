using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Quota/update.json
// this example is just showing the usage of "Quotas_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
AzureLocation location = new AzureLocation("eastus");
MachineLearningQuotaUpdateContent content = new MachineLearningQuotaUpdateContent()
{
    Value =
    {
    new MachineLearningQuotaProperties()
    {
    Id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs",
    QuotaBasePropertiesType = "Microsoft.MachineLearningServices/workspaces/quotas",
    Limit = 100,
    Unit = MachineLearningQuotaUnit.Count,
    },new MachineLearningQuotaProperties()
    {
    Id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs",
    QuotaBasePropertiesType = "Microsoft.MachineLearningServices/workspaces/quotas",
    Limit = 200,
    Unit = MachineLearningQuotaUnit.Count,
    }
    },
};
await foreach (MachineLearningWorkspaceQuotaUpdate item in subscriptionResource.UpdateMachineLearningQuotasAsync(location, content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
