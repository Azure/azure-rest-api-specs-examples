using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/VirtualEnclave_HandleApprovalCreation.json
// this example is just showing the usage of "VirtualEnclave_HandleApprovalCreation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveResource created on azure
// for more information of creating VirtualEnclaveResource, please refer to the document of VirtualEnclaveResource
string subscriptionId = "c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c";
string resourceGroupName = "rgopenapi";
string virtualEnclaveName = "TestMyEnclave";
ResourceIdentifier virtualEnclaveResourceId = VirtualEnclaveResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualEnclaveName);
VirtualEnclaveResource virtualEnclave = client.GetVirtualEnclaveResource(virtualEnclaveResourceId);

// invoke the operation
ApprovalCallbackContent content = new ApprovalCallbackContent(PostActionResourceRequestAction.Create, PostActionCallbackApprovalStatus.Approved)
{
    ApprovalCallbackPayload = "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}",
};
ArmOperation<ApprovalActionResult> lro = await virtualEnclave.HandleApprovalCreationAsync(WaitUntil.Completed, content);
ApprovalActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
