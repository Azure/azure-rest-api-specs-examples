using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/CreateOrUpdateFunctionAppFlexConsumptionWithDetails.json
// this example is just showing the usage of "WebApps_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this WebSiteResource
WebSiteCollection collection = resourceGroupResource.GetWebSites();

// invoke the operation
string name = "sitef6141";
WebSiteData data = new WebSiteData(new AzureLocation("East US"))
{
    SiteConfig = new SiteConfigProperties()
    {
        AppSettings =
        {
        new AppServiceNameValuePair()
        {
        Name = "AzureWebJobsStorage",
        Value = "DefaultEndpointsProtocol=https;AccountName=StorageAccountName;AccountKey=Sanitized;EndpointSuffix=core.windows.net",
        },new AppServiceNameValuePair()
        {
        Name = "APPLICATIONINSIGHTS_CONNECTION_STRING",
        Value = "InstrumentationKey=Sanitized;IngestionEndpoint=Sanitized;LiveEndpoint=Sanitized",
        }
        },
    },
    FunctionAppConfig = new FunctionAppConfig()
    {
        DeploymentStorage = new FunctionAppStorage()
        {
            StorageType = FunctionAppStorageType.BlobContainer,
            Value = new Uri("https://storageAccountName.blob.core.windows.net/containername"),
            Authentication = new FunctionAppStorageAuthentication()
            {
                AuthenticationType = FunctionAppStorageAccountAuthenticationType.StorageAccountConnectionString,
                StorageAccountConnectionStringName = "TheAppSettingName",
            },
        },
        Runtime = new FunctionAppRuntime()
        {
            Name = FunctionAppRuntimeName.Python,
            Version = "3.11",
        },
        ScaleAndConcurrency = new FunctionAppScaleAndConcurrency()
        {
            AlwaysReady =
            {
            new FunctionAppAlwaysReadyConfig()
            {
            Name = "http",
            InstanceCount = 2,
            }
            },
            MaximumInstanceCount = 100,
            InstanceMemoryMB = 2048,
            HttpPerInstanceConcurrency = 16,
        },
    },
    Kind = "functionapp,linux",
};
ArmOperation<WebSiteResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
WebSiteResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebSiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
