using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/SharePut.json
// this example is just showing the usage of "Shares_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeShareResource created on azure
// for more information of creating DataBoxEdgeShareResource, please refer to the document of DataBoxEdgeShareResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
string name = "smbshare";
ResourceIdentifier dataBoxEdgeShareResourceId = DataBoxEdgeShareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName, name);
DataBoxEdgeShareResource dataBoxEdgeShare = client.GetDataBoxEdgeShareResource(dataBoxEdgeShareResourceId);

// invoke the operation
DataBoxEdgeShareData data = new DataBoxEdgeShareData(new ShareStatus("Online"), DataBoxEdgeShareMonitoringStatus.Enabled, ShareAccessProtocol.Smb)
{
    Description = "",
    AzureContainerInfo = new DataBoxEdgeStorageContainerInfo(new ResourceIdentifier("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/sac1"), "testContainerSMB", DataBoxEdgeStorageContainerDataFormat.BlockBlob),
    UserAccessRights = { new UserAccessRight(new ResourceIdentifier("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/users/user2"), ShareAccessType.Change) },
    DataPolicy = DataBoxEdgeDataPolicy.Cloud,
};
ArmOperation<DataBoxEdgeShareResource> lro = await dataBoxEdgeShare.UpdateAsync(WaitUntil.Completed, data);
DataBoxEdgeShareResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxEdgeShareData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
