using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationFabrics_ReassociateGateway.json
// this example is just showing the usage of "ReplicationFabrics_ReassociateGateway" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryFabricResource created on azure
// for more information of creating SiteRecoveryFabricResource, please refer to the document of SiteRecoveryFabricResource
string subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
string resourceGroupName = "MadhaviVRG";
string resourceName = "MadhaviVault";
string fabricName = "GRACE-V2A-1";
ResourceIdentifier siteRecoveryFabricResourceId = SiteRecoveryFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName);
SiteRecoveryFabricResource siteRecoveryFabric = client.GetSiteRecoveryFabricResource(siteRecoveryFabricResourceId);

// invoke the operation
FailoverProcessServerContent content = new FailoverProcessServerContent()
{
    Properties = new FailoverProcessServerProperties()
    {
        ContainerName = "cloud_1f3c15af-2256-4568-9e06-e1ef4f728f75",
        SourceProcessServerId = Guid.Parse("AFA0EC54-1894-4E44-9CAB02DB8854B117"),
        TargetProcessServerId = Guid.Parse("5D3ED340-85AE-C646-B338641E015DA405"),
        VmsToMigrate =
        {
        "Vm1","Vm2"
        },
        UpdateType = "ServerLevel",
    },
};
ArmOperation<SiteRecoveryFabricResource> lro = await siteRecoveryFabric.ReassociateGatewayAsync(WaitUntil.Completed, content);
SiteRecoveryFabricResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryFabricData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
