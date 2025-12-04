using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkVirtualApplianceBootDiagnostics.json
// this example is just showing the usage of "NetworkVirtualAppliances_GetBootDiagnosticLogs" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkVirtualApplianceResource created on azure
// for more information of creating NetworkVirtualApplianceResource, please refer to the document of NetworkVirtualApplianceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkVirtualApplianceName = "nva";
ResourceIdentifier networkVirtualApplianceResourceId = NetworkVirtualApplianceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkVirtualApplianceName);
NetworkVirtualApplianceResource networkVirtualAppliance = client.GetNetworkVirtualApplianceResource(networkVirtualApplianceResourceId);

// invoke the operation
NetworkVirtualApplianceBootDiagnosticContent content = new NetworkVirtualApplianceBootDiagnosticContent
{
    InstanceId = 0,
    SerialConsoleStorageSasUri = new Uri("https://blobcortextesturl.blob.core.windows.net/nvaBootDiagContainer/serialLogs.txt?sp=rw&se=2018-01-10T03%3A42%3A04Z&sv=2017-04-17&sig=WvXrT5bDmDFfgHs%2Brz%2BjAu123eRCNE9BO0eQYcPDT7pY%3D&sr=b"),
    ConsoleScreenshotStorageSasUri = new Uri("https://blobcortextesturl.blob.core.windows.net/nvaBootDiagContainer/consoleScreenshot.png?sp=rw&se=2018-01-10T03%3A42%3A04Z&sv=2017-04-17&sig=WvXrT5bDmDFfgHs%2Brz%2BjAu123eRCNE9BO0eQYcPDT7pY%3D&sr=b"),
};
ArmOperation<NetworkVirtualApplianceInstanceId> lro = await networkVirtualAppliance.GetBootDiagnosticLogsAsync(WaitUntil.Completed, content);
NetworkVirtualApplianceInstanceId result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
