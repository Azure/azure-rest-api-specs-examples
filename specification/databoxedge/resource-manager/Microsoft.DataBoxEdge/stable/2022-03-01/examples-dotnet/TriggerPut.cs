using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/TriggerPut.json
// this example is just showing the usage of "Triggers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeTriggerResource created on azure
// for more information of creating DataBoxEdgeTriggerResource, please refer to the document of DataBoxEdgeTriggerResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
string name = "trigger1";
ResourceIdentifier dataBoxEdgeTriggerResourceId = DataBoxEdgeTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName, name);
DataBoxEdgeTriggerResource dataBoxEdgeTrigger = client.GetDataBoxEdgeTriggerResource(dataBoxEdgeTriggerResourceId);

// invoke the operation
DataBoxEdgeTriggerData data = new EdgeFileEventTrigger(new EdgeFileSourceInfo(new ResourceIdentifier("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/shares/share1")), new DataBoxEdgeRoleSinkInfo(new ResourceIdentifier("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/role1")))
{
    CustomContextTag = "CustomContextTags-1235346475",
};
ArmOperation<DataBoxEdgeTriggerResource> lro = await dataBoxEdgeTrigger.UpdateAsync(WaitUntil.Completed, data);
DataBoxEdgeTriggerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxEdgeTriggerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
