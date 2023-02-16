using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Advisor;
using Azure.ResourceManager.Advisor.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
// this example is just showing the usage of "Configurations_CreateInSubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subscriptionId";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
ConfigurationName configurationName = ConfigurationName.Default;
ConfigData data = new ConfigData()
{
    Exclude = true,
    LowCpuThreshold = CpuThreshold.Five,
    Digests =
    {
    new DigestConfig()
    {
    Name = "digestConfigName",
    ActionGroupResourceId = "/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName",
    Frequency = 30,
    Categories =
    {
    Category.HighAvailability,Category.Security,Category.Performance,Category.Cost,Category.OperationalExcellence
    },
    Language = "en",
    State = DigestConfigState.Active,
    }
    },
};
ConfigData result = await subscriptionResource.CreateConfigurationAsync(configurationName, data);

Console.WriteLine($"Succeeded: {result}");
