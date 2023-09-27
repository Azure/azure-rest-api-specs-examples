using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DatabaseFleetManager;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-06-15-preview/examples/FleetMembers_Create.json
// this example is just showing the usage of "FleetMembers_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseFleetResource created on azure
// for more information of creating DatabaseFleetResource, please refer to the document of DatabaseFleetResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
ResourceIdentifier databaseFleetResourceId = DatabaseFleetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName);
DatabaseFleetResource databaseFleet = client.GetDatabaseFleetResource(databaseFleetResourceId);

// get the collection of this DatabaseFleetMemberResource
DatabaseFleetMemberCollection collection = databaseFleet.GetDatabaseFleetMembers();

// invoke the operation
string fleetMemberName = "member-1";
DatabaseFleetMemberData data = new DatabaseFleetMemberData()
{
    ClusterResourceId = new ResourceIdentifier("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/cluster-1"),
};
ArmOperation<DatabaseFleetMemberResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, fleetMemberName, data);
DatabaseFleetMemberResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseFleetMemberData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
