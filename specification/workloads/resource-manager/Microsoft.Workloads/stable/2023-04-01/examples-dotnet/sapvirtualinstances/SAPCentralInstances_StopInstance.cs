using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Workloads.Models;
using Azure.ResourceManager.Workloads;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPCentralInstances_StopInstance.json
// this example is just showing the usage of "SAPCentralInstances_StopInstance" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapCentralServerInstanceResource created on azure
// for more information of creating SapCentralServerInstanceResource, please refer to the document of SapCentralServerInstanceResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
string resourceGroupName = "test-rg";
string sapVirtualInstanceName = "X00";
string centralInstanceName = "centralServer";
ResourceIdentifier sapCentralServerInstanceResourceId = SapCentralServerInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapVirtualInstanceName, centralInstanceName);
SapCentralServerInstanceResource sapCentralServerInstance = client.GetSapCentralServerInstanceResource(sapCentralServerInstanceResourceId);

// invoke the operation
SapStopContent content = new SapStopContent
{
    SoftStopTimeoutSeconds = 1200L,
};
ArmOperation<OperationStatusResult> lro = await sapCentralServerInstance.StopInstanceAsync(WaitUntil.Completed, content: content);
OperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
