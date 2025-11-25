using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Datadog.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Datadog;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2025-06-11/examples/Monitors_Create.json
// this example is just showing the usage of "Monitors_Create" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this DatadogMonitorResource
DatadogMonitorCollection collection = resourceGroupResource.GetDatadogMonitors();

// invoke the operation
string monitorName = "myMonitor";
DatadogMonitorData data = new DatadogMonitorData(new AzureLocation("West US"))
{
    Properties = new DatadogMonitorProperties
    {
        MonitoringStatus = DatadogMonitoringStatus.Enabled,
        DatadogOrganizationProperties = new DatadogOrganizationProperties
        {
            Name = "myOrg",
            Id = "myOrg123",
            LinkingAuthCode = "someAuthCode",
            LinkingClientId = "00000000-0000-0000-0000-000000000000",
            EnterpriseAppId = "00000000-0000-0000-0000-000000000000",
            IsCspm = false,
            IsResourceCollection = false,
        },
        UserInfo = new DatadogUserInfo
        {
            Name = "Alice",
            EmailAddress = "alice@microsoft.com",
            PhoneNumber = "123-456-7890",
        },
    },
    SkuName = "free_Monthly",
    Tags =
    {
    ["Environment"] = "Dev"
    },
};
ArmOperation<DatadogMonitorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, monitorName, data);
DatadogMonitorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatadogMonitorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
