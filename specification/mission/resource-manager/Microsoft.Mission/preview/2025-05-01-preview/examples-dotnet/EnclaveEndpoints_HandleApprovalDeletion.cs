using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/EnclaveEndpoints_HandleApprovalDeletion.json
// this example is just showing the usage of "EnclaveEndpoints_HandleApprovalDeletion" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveEndpointResource created on azure
// for more information of creating VirtualEnclaveEndpointResource, please refer to the document of VirtualEnclaveEndpointResource
string subscriptionId = "73CEECEF-2C30-488E-946F-D20F414D99BA";
string resourceGroupName = "rgopenapi";
string virtualEnclaveName = "TestMyEnclave";
string enclaveEndpointName = "TestMyEnclaveEndpoint";
ResourceIdentifier virtualEnclaveEndpointResourceId = VirtualEnclaveEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualEnclaveName, enclaveEndpointName);
VirtualEnclaveEndpointResource virtualEnclaveEndpoint = client.GetVirtualEnclaveEndpointResource(virtualEnclaveEndpointResourceId);

// invoke the operation
ApprovalDeletionCallbackContent content = new ApprovalDeletionCallbackContent(PostActionDeletionResourceRequestAction.Create);
ArmOperation<ApprovalActionResult> lro = await virtualEnclaveEndpoint.HandleApprovalDeletionAsync(WaitUntil.Completed, content);
ApprovalActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
