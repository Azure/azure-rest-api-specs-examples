using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.OperationalInsights;
using Azure.ResourceManager.OperationalInsights.Models;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataSourcesCreate.json
// this example is just showing the usage of "DataSources_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsDataSourceResource created on azure
// for more information of creating OperationalInsightsDataSourceResource, please refer to the document of OperationalInsightsDataSourceResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "OIAutoRest5123";
string workspaceName = "AzTest9724";
string dataSourceName = "AzTestDS774";
ResourceIdentifier operationalInsightsDataSourceResourceId = OperationalInsightsDataSourceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, dataSourceName);
OperationalInsightsDataSourceResource operationalInsightsDataSource = client.GetOperationalInsightsDataSourceResource(operationalInsightsDataSourceResourceId);

// invoke the operation
OperationalInsightsDataSourceData data = new OperationalInsightsDataSourceData(BinaryData.FromObjectAsJson(new Dictionary<string, object>()
{
    ["LinkedResourceId"] = "/subscriptions/00000000-0000-0000-0000-00000000000/providers/microsoft.insights/eventtypes/management"
}), OperationalInsightsDataSourceKind.AzureActivityLog);
ArmOperation<OperationalInsightsDataSourceResource> lro = await operationalInsightsDataSource.UpdateAsync(WaitUntil.Completed, data);
OperationalInsightsDataSourceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OperationalInsightsDataSourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
