using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestDelete.json
// this example is just showing the usage of "WebTests_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationInsightsWebTestResource created on azure
// for more information of creating ApplicationInsightsWebTestResource, please refer to the document of ApplicationInsightsWebTestResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
string webTestName = "my-webtest-01-mywebservice";
ResourceIdentifier applicationInsightsWebTestResourceId = ApplicationInsightsWebTestResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, webTestName);
ApplicationInsightsWebTestResource applicationInsightsWebTest = client.GetApplicationInsightsWebTestResource(applicationInsightsWebTestResourceId);

// invoke the operation
await applicationInsightsWebTest.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
