using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsCreate.json
// this example is just showing the usage of "Components_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ApplicationInsightsComponentResource
ApplicationInsightsComponentCollection collection = resourceGroupResource.GetApplicationInsightsComponents();

// invoke the operation
string resourceName = "my-component";
ApplicationInsightsComponentData data = new ApplicationInsightsComponentData(new AzureLocation("South Central US"), "web")
{
    ApplicationType = ApplicationInsightsApplicationType.Web,
    FlowType = ComponentFlowType.Bluefield,
    RequestSource = ComponentRequestSource.Rest,
    WorkspaceResourceId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace"),
};
ArmOperation<ApplicationInsightsComponentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
ApplicationInsightsComponentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApplicationInsightsComponentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
