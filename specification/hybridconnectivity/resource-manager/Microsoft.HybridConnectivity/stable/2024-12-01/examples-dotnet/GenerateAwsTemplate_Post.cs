using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: 2024-12-01/GenerateAwsTemplate_Post.json
// this example is just showing the usage of "GenerateAwsTemplate_Post" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "5ACC4579-DB34-4C2F-8F8C-25061168F342";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
GenerateAwsTemplateContent content = new GenerateAwsTemplateContent(new ResourceIdentifier("pnxcfjidglabnwxit"))
{
    SolutionTypes = {new PublicCloudConnectorSolutionTypeSettings("hjyownzpfxwiufmd")
    {
    SolutionSettings = new PublicCloudConnectorSolutionSettings(),
    }},
};
GenerateAwsTemplateResult result = await subscriptionResource.PostGenerateAwsTemplateAsync(content);

Console.WriteLine($"Succeeded: {result}");
