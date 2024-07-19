using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestCreateStandard.json
// this example is just showing the usage of "WebTests_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ApplicationInsightsWebTestResource
ApplicationInsightsWebTestCollection collection = resourceGroupResource.GetApplicationInsightsWebTests();

// invoke the operation
string webTestName = "my-webtest-my-component";
ApplicationInsightsWebTestData data = new ApplicationInsightsWebTestData(new AzureLocation("South Central US"))
{
    SyntheticMonitorId = "my-webtest-my-component",
    WebTestName = "my-webtest-my-component",
    Description = "Ping web test alert for mytestwebapp",
    IsEnabled = true,
    FrequencyInSeconds = 900,
    TimeoutInSeconds = 120,
    WebTestKind = WebTestKind.Standard,
    IsRetryEnabled = true,
    Locations =
    {
    new WebTestGeolocation()
    {
    Location = new AzureLocation("us-fl-mia-edge"),
    }
    },
    Request = new WebTestRequest()
    {
        RequestUri = new Uri("https://bing.com"),
        Headers =
        {
        new WebTestRequestHeaderField()
        {
        HeaderFieldName = "Content-Language",
        HeaderFieldValue = "de-DE",
        },new WebTestRequestHeaderField()
        {
        HeaderFieldName = "Accept-Language",
        HeaderFieldValue = "de-DE",
        }
        },
        HttpVerb = "POST",
        RequestBody = "SGVsbG8gd29ybGQ=",
    },
    ValidationRules = new WebTestValidationRules()
    {
        CheckSsl = true,
        SslCertRemainingLifetimeCheck = 100,
    },
};
ArmOperation<ApplicationInsightsWebTestResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, webTestName, data);
ApplicationInsightsWebTestResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApplicationInsightsWebTestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
