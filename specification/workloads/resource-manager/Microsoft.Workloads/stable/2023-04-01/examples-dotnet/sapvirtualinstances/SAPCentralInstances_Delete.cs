using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Workloads.Models;
using Azure.ResourceManager.Workloads;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPCentralInstances_Delete.json
// this example is just showing the usage of "SAPCentralInstances_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapCentralServerInstanceResource created on azure
// for more information of creating SapCentralServerInstanceResource, please refer to the document of SapCentralServerInstanceResource
string subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
string resourceGroupName = "test-rg";
string sapVirtualInstanceName = "X00";
string centralInstanceName = "centralServer";
ResourceIdentifier sapCentralServerInstanceResourceId = SapCentralServerInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapVirtualInstanceName, centralInstanceName);
SapCentralServerInstanceResource sapCentralServerInstance = client.GetSapCentralServerInstanceResource(sapCentralServerInstanceResourceId);

// invoke the operation
ArmOperation<OperationStatusResult> lro = await sapCentralServerInstance.DeleteAsync(WaitUntil.Completed);
OperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
