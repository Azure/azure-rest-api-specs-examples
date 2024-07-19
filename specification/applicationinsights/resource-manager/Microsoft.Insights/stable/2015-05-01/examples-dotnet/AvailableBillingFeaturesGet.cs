using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AvailableBillingFeaturesGet.json
// this example is just showing the usage of "ComponentAvailableFeatures_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationInsightsComponentResource created on azure
// for more information of creating ApplicationInsightsComponentResource, please refer to the document of ApplicationInsightsComponentResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
string resourceName = "my-component";
ResourceIdentifier applicationInsightsComponentResourceId = ApplicationInsightsComponentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ApplicationInsightsComponentResource applicationInsightsComponent = client.GetApplicationInsightsComponentResource(applicationInsightsComponentResourceId);

// invoke the operation
ApplicationInsightsComponentAvailableFeatures result = await applicationInsightsComponent.GetComponentAvailableFeatureAsync();

Console.WriteLine($"Succeeded: {result}");
