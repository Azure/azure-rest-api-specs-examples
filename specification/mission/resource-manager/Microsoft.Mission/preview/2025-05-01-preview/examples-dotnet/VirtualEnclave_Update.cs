using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/VirtualEnclave_Update.json
// this example is just showing the usage of "EnclaveResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveResource created on azure
// for more information of creating VirtualEnclaveResource, please refer to the document of VirtualEnclaveResource
string subscriptionId = "CA1CB369-DD26-4DB2-9D43-9AFEF0F22093";
string resourceGroupName = "rgopenapi";
string virtualEnclaveName = "TestMyEnclave";
ResourceIdentifier virtualEnclaveResourceId = VirtualEnclaveResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualEnclaveName);
VirtualEnclaveResource virtualEnclave = client.GetVirtualEnclaveResource(virtualEnclaveResourceId);

// invoke the operation
VirtualEnclavePatch patch = new VirtualEnclavePatch
{
    Properties = new VirtualEnclavePatchProperties(new EnclaveVirtualNetwork
    {
        NetworkSize = "small",
        CustomCidrRange = "10.0.0.0/24",
        SubnetConfigurations = { new VirtualEnclaveSubnetConfiguration("test", 26) },
        AllowSubnetCommunication = true,
    })
    {
        IsBastionEnabled = true,
        EnclaveRoleAssignments = {new VirtualEnclaveRoleAssignmentItem("b24988ac-6180-42a0-ab88-20f7382dd24c")
        {
        Principals = {new VirtualEnclavePrincipal("355a6bb0-abc0-4cba-000d-12a345b678c9", VirtualEnclavePrincipalType.User), new VirtualEnclavePrincipal("355a6bb0-abc0-4cba-000d-12a345b678c0", VirtualEnclavePrincipalType.User)},
        }, new VirtualEnclaveRoleAssignmentItem("18d7d88d-d35e-4fb5-a5c3-7773c20a72d9")
        {
        Principals = {new VirtualEnclavePrincipal("355a6bb0-abc0-4cba-000d-12a345b678c9", VirtualEnclavePrincipalType.User)},
        }},
        WorkloadRoleAssignments = {new VirtualEnclaveRoleAssignmentItem("d73bb868-a0df-4d4d-bd69-98a00b01fccb")
        {
        Principals = {new VirtualEnclavePrincipal("01234567-89ab-ef01-2345-0123456789ab", VirtualEnclavePrincipalType.Group)},
        }, new VirtualEnclaveRoleAssignmentItem("fb879df8-f326-4884-b1cf-06f3ad86be52")
        {
        Principals = {new VirtualEnclavePrincipal("01234567-89ab-ef01-2345-0123456789ab", VirtualEnclavePrincipalType.Group)},
        }},
        EnclaveDefaultDiagnosticDestination = VirtualEnclaveDiagnosticDestination.Both,
        MaintenanceModeConfiguration = new VirtualEnclaveMaintenanceModeConfigurationPatch(VirtualEnclaveMaintenanceMode.Off)
        {
            Principals = { new VirtualEnclavePrincipal("355a6bb0-abc0-4cba-000d-12a345b678c9", VirtualEnclavePrincipalType.User) },
            Justification = VirtualEnclaveMaintenanceJustification.Off,
        },
    },
    Tags =
    {
    ["Tag1"] = "Value1"
    },
};
ArmOperation<VirtualEnclaveResource> lro = await virtualEnclave.UpdateAsync(WaitUntil.Completed, patch);
VirtualEnclaveResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
