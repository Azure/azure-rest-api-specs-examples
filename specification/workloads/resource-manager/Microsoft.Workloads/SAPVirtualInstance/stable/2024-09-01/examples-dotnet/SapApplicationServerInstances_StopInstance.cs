using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadsSapVirtualInstance.Models;
using Azure.ResourceManager.WorkloadsSapVirtualInstance;

// Generated from example definition: 2024-09-01/SapApplicationServerInstances_StopInstance.json
// this example is just showing the usage of "SAPApplicationServerInstances_Stop" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapApplicationServerInstanceResource created on azure
// for more information of creating SapApplicationServerInstanceResource, please refer to the document of SapApplicationServerInstanceResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
string resourceGroupName = "test-rg";
string sapVirtualInstanceName = "X00";
string applicationInstanceName = "app01";
ResourceIdentifier sapApplicationServerInstanceResourceId = SapApplicationServerInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapVirtualInstanceName, applicationInstanceName);
SapApplicationServerInstanceResource sapApplicationServerInstance = client.GetSapApplicationServerInstanceResource(sapApplicationServerInstanceResourceId);

// invoke the operation
StopSapInstanceContent content = new StopSapInstanceContent
{
    SoftStopTimeoutSeconds = 0L,
};
ArmOperation<OperationStatusResult> lro = await sapApplicationServerInstance.StopAsync(WaitUntil.Completed, content: content);
OperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
