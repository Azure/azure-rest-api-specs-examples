using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestUpdateTagsOnly.json
// this example is just showing the usage of "WebTests_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebTestResource created on azure
// for more information of creating WebTestResource, please refer to the document of WebTestResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
string webTestName = "my-webtest-my-component";
ResourceIdentifier webTestResourceId = WebTestResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, webTestName);
WebTestResource webTest = client.GetWebTestResource(webTestResourceId);

// invoke the operation
ComponentTag webTestTags = new ComponentTag()
{
    Tags =
    {
    ["Color"] = "AzureBlue",
    ["CustomField-01"] = "This is a random value",
    ["SystemType"] = "A08",
    ["hidden-link:/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component"] = "Resource",
    },
};
WebTestResource result = await webTest.UpdateAsync(webTestTags);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebTestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
