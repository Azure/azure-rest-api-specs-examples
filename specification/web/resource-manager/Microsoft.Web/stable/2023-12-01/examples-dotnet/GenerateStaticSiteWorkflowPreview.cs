using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GenerateStaticSiteWorkflowPreview.json
// this example is just showing the usage of "StaticSites_PreviewWorkflow" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("West US 2");
StaticSitesWorkflowPreviewContent content = new StaticSitesWorkflowPreviewContent()
{
    RepositoryUri = new Uri("https://github.com/username/RepoName"),
    Branch = "master",
    BuildProperties = new StaticSiteBuildProperties()
    {
        AppLocation = "app",
        ApiLocation = "api",
        AppArtifactLocation = "build",
    },
};
StaticSitesWorkflowPreview result = await subscriptionResource.PreviewStaticSiteWorkflowAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
