using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/InstanceFailoverGroupCreateOrUpdate.json
// this example is just showing the usage of "InstanceFailoverGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InstanceFailoverGroupResource created on azure
// for more information of creating InstanceFailoverGroupResource, please refer to the document of InstanceFailoverGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
AzureLocation locationName = new AzureLocation("Japan East");
string failoverGroupName = "failover-group-test-3";
ResourceIdentifier instanceFailoverGroupResourceId = InstanceFailoverGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, locationName, failoverGroupName);
InstanceFailoverGroupResource instanceFailoverGroup = client.GetInstanceFailoverGroupResource(instanceFailoverGroupResourceId);

// invoke the operation
InstanceFailoverGroupData data = new InstanceFailoverGroupData
{
    SecondaryType = GeoSecondaryInstanceType.Geo,
    ReadWriteEndpoint = new InstanceFailoverGroupReadWriteEndpoint(ReadWriteEndpointFailoverPolicy.Automatic)
    {
        FailoverWithDataLossGracePeriodMinutes = 480,
    },
    ReadOnlyEndpointFailoverPolicy = ReadOnlyEndpointFailoverPolicy.Disabled,
    PartnerRegions = {new PartnerRegionInfo
    {
    Location = new AzureLocation("Japan West"),
    }},
    ManagedInstancePairs = {new ManagedInstancePairInfo
    {
    PrimaryManagedInstanceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-primary-mngdInstance"),
    PartnerManagedInstanceId = new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-secondary-mngdInstance"),
    }},
};
ArmOperation<InstanceFailoverGroupResource> lro = await instanceFailoverGroup.UpdateAsync(WaitUntil.Completed, data);
InstanceFailoverGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InstanceFailoverGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
