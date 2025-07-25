using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Grafana;

// Generated from example definition: 2024-11-01-preview/Dashboard_Create.json
// this example is just showing the usage of "ManagedDashboard_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ManagedDashboardResource
ManagedDashboardCollection collection = resourceGroupResource.GetManagedDashboards();

// invoke the operation
string dashboardName = "myDashboard";
ManagedDashboardData data = new ManagedDashboardData(new AzureLocation("West US"))
{
    Tags =
    {
    ["Environment"] = "Dev"
    },
};
ArmOperation<ManagedDashboardResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dashboardName, data);
ManagedDashboardResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedDashboardData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
