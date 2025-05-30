using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DigitalTwins.Models;
using Azure.ResourceManager.DigitalTwins;

// Generated from example definition: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/TimeSeriesDatabaseConnectionsDelete_example.json
// this example is just showing the usage of "TimeSeriesDatabaseConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TimeSeriesDatabaseConnectionResource created on azure
// for more information of creating TimeSeriesDatabaseConnectionResource, please refer to the document of TimeSeriesDatabaseConnectionResource
string subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
string resourceGroupName = "resRg";
string resourceName = "myDigitalTwinsService";
string timeSeriesDatabaseConnectionName = "myConnection";
ResourceIdentifier timeSeriesDatabaseConnectionResourceId = TimeSeriesDatabaseConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, timeSeriesDatabaseConnectionName);
TimeSeriesDatabaseConnectionResource timeSeriesDatabaseConnection = client.GetTimeSeriesDatabaseConnectionResource(timeSeriesDatabaseConnectionResourceId);

// invoke the operation
ArmOperation<TimeSeriesDatabaseConnectionResource> lro = await timeSeriesDatabaseConnection.DeleteAsync(WaitUntil.Completed);
TimeSeriesDatabaseConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TimeSeriesDatabaseConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
