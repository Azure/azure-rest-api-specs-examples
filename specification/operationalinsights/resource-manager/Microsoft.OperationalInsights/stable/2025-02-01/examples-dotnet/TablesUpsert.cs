using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2025-02-01/examples/TablesUpsert.json
// this example is just showing the usage of "Tables_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsWorkspaceResource created on azure
// for more information of creating OperationalInsightsWorkspaceResource, please refer to the document of OperationalInsightsWorkspaceResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "oiautorest6685";
string workspaceName = "oiautorest6685";
ResourceIdentifier operationalInsightsWorkspaceResourceId = OperationalInsightsWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
OperationalInsightsWorkspaceResource operationalInsightsWorkspace = client.GetOperationalInsightsWorkspaceResource(operationalInsightsWorkspaceResourceId);

// get the collection of this OperationalInsightsTableResource
OperationalInsightsTableCollection collection = operationalInsightsWorkspace.GetOperationalInsightsTables();

// invoke the operation
string tableName = "AzureNetworkFlow";
OperationalInsightsTableData data = new OperationalInsightsTableData
{
    RetentionInDays = 45,
    TotalRetentionInDays = 70,
    Schema = new OperationalInsightsSchema
    {
        Name = "AzureNetworkFlow",
        Columns = {new OperationalInsightsColumn
        {
        Name = "MyNewColumn",
        ColumnType = OperationalInsightsColumnType.Guid,
        }},
    },
};
ArmOperation<OperationalInsightsTableResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, tableName, data);
OperationalInsightsTableResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OperationalInsightsTableData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
