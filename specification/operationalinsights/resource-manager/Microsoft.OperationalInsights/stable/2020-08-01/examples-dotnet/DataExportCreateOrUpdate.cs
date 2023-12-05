using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportCreateOrUpdate.json
// this example is just showing the usage of "DataExports_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsWorkspaceResource created on azure
// for more information of creating OperationalInsightsWorkspaceResource, please refer to the document of OperationalInsightsWorkspaceResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "RgTest1";
string workspaceName = "DeWnTest1234";
ResourceIdentifier operationalInsightsWorkspaceResourceId = OperationalInsightsWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
OperationalInsightsWorkspaceResource operationalInsightsWorkspace = client.GetOperationalInsightsWorkspaceResource(operationalInsightsWorkspaceResourceId);

// get the collection of this OperationalInsightsDataExportResource
OperationalInsightsDataExportCollection collection = operationalInsightsWorkspace.GetOperationalInsightsDataExports();

// invoke the operation
string dataExportName = "export1";
OperationalInsightsDataExportData data = new OperationalInsightsDataExportData()
{
    TableNames =
    {
    "Heartbeat"
    },
    ResourceId = new ResourceIdentifier("/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test"),
};
ArmOperation<OperationalInsightsDataExportResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataExportName, data);
OperationalInsightsDataExportResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OperationalInsightsDataExportData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
