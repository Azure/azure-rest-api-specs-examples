using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotFirmwareDefense;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2025-08-02/examples/UsageMetrics_Get_MaximumSet_Gen.json
// this example is just showing the usage of "UsageMetrics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirmwareAnalysisWorkspaceResource created on azure
// for more information of creating FirmwareAnalysisWorkspaceResource, please refer to the document of FirmwareAnalysisWorkspaceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgiotfirmwaredefense";
string workspaceName = "exampleWorkspaceName";
ResourceIdentifier firmwareAnalysisWorkspaceResourceId = FirmwareAnalysisWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
FirmwareAnalysisWorkspaceResource firmwareAnalysisWorkspace = client.GetFirmwareAnalysisWorkspaceResource(firmwareAnalysisWorkspaceResourceId);

// get the collection of this UsageMetricResource
UsageMetricCollection collection = firmwareAnalysisWorkspace.GetUsageMetrics();

// invoke the operation
string name = "default";
NullableResponse<UsageMetricResource> response = await collection.GetIfExistsAsync(name);
UsageMetricResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    UsageMetricData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
