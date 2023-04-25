using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPApplicationServerInstances_StartInstance.json
// this example is just showing the usage of "SAPApplicationServerInstances_StartInstance" operation, for the dependent resources, they will have to be created separately.

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
ArmOperation<OperationStatusResult> lro = await sapApplicationServerInstance.StartInstanceAsync(WaitUntil.Completed);
OperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
