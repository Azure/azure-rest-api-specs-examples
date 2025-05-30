using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/ProjectEnvironmentTypes_Put.json
// this example is just showing the usage of "ProjectEnvironmentTypes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevCenterProjectResource created on azure
// for more information of creating DevCenterProjectResource, please refer to the document of DevCenterProjectResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
string projectName = "ContosoProj";
ResourceIdentifier devCenterProjectResourceId = DevCenterProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName);
DevCenterProjectResource devCenterProject = client.GetDevCenterProjectResource(devCenterProjectResourceId);

// get the collection of this DevCenterProjectEnvironmentResource
DevCenterProjectEnvironmentCollection collection = devCenterProject.GetDevCenterProjectEnvironments();

// invoke the operation
string environmentTypeName = "DevTest";
DevCenterProjectEnvironmentData data = new DevCenterProjectEnvironmentData(default)
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1")] = new UserAssignedIdentity()
        },
    },
    DeploymentTargetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000"),
    Status = EnvironmentTypeEnableStatus.IsEnabled,
    Roles =
    {
    ["4cbf0b6c-e750-441c-98a7-10da8387e4d6"] = new DevCenterEnvironmentRole()
    },
    UserRoleAssignments =
    {
    ["e45e3m7c-176e-416a-b466-0c5ec8298f8a"] = new DevCenterUserRoleAssignments
    {
    Roles =
    {
    ["4cbf0b6c-e750-441c-98a7-10da8387e4d6"] = new DevCenterEnvironmentRole()
    },
    }
    },
    Tags =
    {
    ["CostCenter"] = "RnD"
    },
};
ArmOperation<DevCenterProjectEnvironmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, environmentTypeName, data);
DevCenterProjectEnvironmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevCenterProjectEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
