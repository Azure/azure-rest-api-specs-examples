using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/UpdateDiagnosticRemoteSupportSettings.json
// this example is just showing the usage of "DiagnosticSettings_UpdateDiagnosticRemoteSupportSettings" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiagnosticRemoteSupportSettingResource created on azure
// for more information of creating DiagnosticRemoteSupportSettingResource, please refer to the document of DiagnosticRemoteSupportSettingResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
ResourceIdentifier diagnosticRemoteSupportSettingResourceId = DiagnosticRemoteSupportSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName);
DiagnosticRemoteSupportSettingResource diagnosticRemoteSupportSetting = client.GetDiagnosticRemoteSupportSettingResource(diagnosticRemoteSupportSettingResourceId);

// invoke the operation
DiagnosticRemoteSupportSettingData data = new DiagnosticRemoteSupportSettingData
{
    RemoteSupportSettingsList = {new EdgeRemoteSupportSettings
    {
    RemoteApplicationType = EdgeRemoteApplicationType.Powershell,
    AccessLevel = EdgeRemoteApplicationAccessLevel.ReadWrite,
    ExpireOn = DateTimeOffset.Parse("2021-07-07T00:00:00+00:00"),
    }},
};
ArmOperation<DiagnosticRemoteSupportSettingResource> lro = await diagnosticRemoteSupportSetting.CreateOrUpdateAsync(WaitUntil.Completed, data);
DiagnosticRemoteSupportSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiagnosticRemoteSupportSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
