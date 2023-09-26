using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DatabaseFleetManager;
using Azure.ResourceManager.DatabaseFleetManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-06-15-preview/examples/Fleets_Get.json
// this example is just showing the usage of "Fleets_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
DatabaseFleetResource result = await databaseFleet.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseFleetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
