using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/CalculateModelCapacity.json
// this example is just showing the usage of "CalculateModelCapacity" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
CalculateModelCapacityContent content = new CalculateModelCapacityContent()
{
    Model = new CognitiveServicesAccountDeploymentModel()
    {
        Format = "OpenAI",
        Name = "gpt-4",
        Version = "0613",
    },
    SkuName = "ProvisionedManaged",
    Workloads =
    {
    new ModelCapacityCalculatorWorkload()
    {
    RequestPerMinute = 10L,
    RequestParameters = new ModelCapacityCalculatorWorkloadRequestParam()
    {
    AvgPromptTokens = 30L,
    AvgGeneratedTokens = 50L,
    },
    },new ModelCapacityCalculatorWorkload()
    {
    RequestPerMinute = 20L,
    RequestParameters = new ModelCapacityCalculatorWorkloadRequestParam()
    {
    AvgPromptTokens = 60L,
    AvgGeneratedTokens = 20L,
    },
    }
    },
};
CalculateModelCapacityResult result = await subscriptionResource.CalculateModelCapacityAsync(content);

Console.WriteLine($"Succeeded: {result}");
