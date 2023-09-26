using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DatabaseFleetManager;
using Azure.ResourceManager.DatabaseFleetManager.Models;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-06-15-preview/examples/UpdateRuns_Delete.json
// this example is just showing the usage of "UpdateRuns_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseFleetUpdateRunResource created on azure
// for more information of creating DatabaseFleetUpdateRunResource, please refer to the document of DatabaseFleetUpdateRunResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
string updateRunName = "run1";
ResourceIdentifier databaseFleetUpdateRunResourceId = DatabaseFleetUpdateRunResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, updateRunName);
DatabaseFleetUpdateRunResource databaseFleetUpdateRun = client.GetDatabaseFleetUpdateRunResource(databaseFleetUpdateRunResourceId);

// invoke the operation
await databaseFleetUpdateRun.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
