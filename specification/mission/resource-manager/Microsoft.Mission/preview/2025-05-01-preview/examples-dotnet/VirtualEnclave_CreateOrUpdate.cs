using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/VirtualEnclave_CreateOrUpdate.json
// this example is just showing the usage of "EnclaveResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VirtualEnclaveResource
VirtualEnclaveCollection collection = resourceGroupResource.GetVirtualEnclaves();

// invoke the operation
string virtualEnclaveName = "TestMyEnclave";
VirtualEnclaveData data = new VirtualEnclaveData(new AzureLocation("westcentralus"))
{
    Properties = new VirtualEnclaveProperties(new EnclaveVirtualNetwork
    {
        NetworkSize = "small",
        CustomCidrRange = "10.0.0.0/24",
        SubnetConfigurations = { new VirtualEnclaveSubnetConfiguration("test", 26) },
        AllowSubnetCommunication = true,
    }, new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/microsoft.mission/communities/TestMyCommunity"))
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
        EnclaveDefaultSettings = new EnclaveDefaultSettings
        {
            DiagnosticDestination = VirtualEnclaveDiagnosticDestination.Both,
        },
        MaintenanceModeConfiguration = new VirtualEnclaveMaintenanceModeConfiguration(VirtualEnclaveMaintenanceMode.Off)
        {
            Principals = { new VirtualEnclavePrincipal("355a6bb0-abc0-4cba-000d-12a345b678c9", VirtualEnclavePrincipalType.User) },
            Justification = VirtualEnclaveMaintenanceJustification.Off,
        },
    },
    Identity = new ManagedServiceIdentity("SystemAssigned,UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["Tag1"] = "Value1"
    },
};
ArmOperation<VirtualEnclaveResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, virtualEnclaveName, data);
VirtualEnclaveResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
