using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/EnclaveConnection_HandleApprovalCreation.json
// this example is just showing the usage of "EnclaveConnection_HandleApprovalCreation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveConnectionResource created on azure
// for more information of creating VirtualEnclaveConnectionResource, please refer to the document of VirtualEnclaveConnectionResource
string subscriptionId = "73CEECEF-2C30-488E-946F-D20F414D99BA";
string resourceGroupName = "rgopenapi";
string enclaveConnectionName = "TestMyEnclaveConnection";
ResourceIdentifier virtualEnclaveConnectionResourceId = VirtualEnclaveConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, enclaveConnectionName);
VirtualEnclaveConnectionResource virtualEnclaveConnection = client.GetVirtualEnclaveConnectionResource(virtualEnclaveConnectionResourceId);

// invoke the operation
ApprovalCallbackContent content = new ApprovalCallbackContent(PostActionResourceRequestAction.Create, PostActionCallbackApprovalStatus.Approved)
{
    ApprovalCallbackPayload = "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}",
};
ArmOperation<ApprovalActionResult> lro = await virtualEnclaveConnection.HandleApprovalCreationAsync(WaitUntil.Completed, content);
ApprovalActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
