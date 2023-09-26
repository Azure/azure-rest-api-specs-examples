using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DatabaseFleetManager;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-06-15-preview/examples/FleetMembers_ListByFleet.json
// this example is just showing the usage of "FleetMembers_ListByFleet" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (DatabaseFleetMemberResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DatabaseFleetMemberData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
