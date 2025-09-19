using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/Approvals_Update.json
// this example is just showing the usage of "ApprovalResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveApprovalResource created on azure
// for more information of creating VirtualEnclaveApprovalResource, please refer to the document of VirtualEnclaveApprovalResource
string resourceUri = "subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/enclaveconnections/TestMyEnclaveConnection";
string approvalName = "TestApprovals";
ResourceIdentifier virtualEnclaveApprovalResourceId = VirtualEnclaveApprovalResource.CreateResourceIdentifier(resourceUri, approvalName);
VirtualEnclaveApprovalResource virtualEnclaveApproval = client.GetVirtualEnclaveApprovalResource(virtualEnclaveApprovalResourceId);

// invoke the operation
VirtualEnclaveApprovalPatch patch = new VirtualEnclaveApprovalPatch
{
    Properties = new VirtualEnclaveApprovalPatchProperties(new ApprovalRequestMetadataPatch("string")
    {
        ApprovalCallbackRoute = "approvalCallback",
        ApprovalCallbackPayload = "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}",
        ApprovalStatus = VirtualEnclaveApprovalStatus.Approved,
    })
    {
        ParentResourceId = new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/microsoft.mission/virtualenclaves/TestMyEnclave"),
        GrandparentResourceId = new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/testrg/providers/Microsoft.Mission/communities/TestMyCommunity"),
        Approvers = {new VirtualEnclaveApprover("00000000-0000-0000-0000-000000000000", DateTimeOffset.Parse("2023-03-17T20:43:17.760Z"))
        {
        ActionPerformed = ActionPerformed.Approved,
        }},
        TicketId = "string",
        CreatedOn = DateTimeOffset.Parse("2023-03-17T20:43:17.760Z"),
        StateChangedOn = DateTimeOffset.Parse("2023-03-17T20:43:17.760Z"),
    },
};
ArmOperation<VirtualEnclaveApprovalResource> lro = await virtualEnclaveApproval.UpdateAsync(WaitUntil.Completed, patch);
VirtualEnclaveApprovalResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveApprovalData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
