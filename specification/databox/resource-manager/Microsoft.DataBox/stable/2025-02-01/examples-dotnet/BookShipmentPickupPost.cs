using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBox.Models;
using Azure.ResourceManager.DataBox;

// Generated from example definition: specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/BookShipmentPickupPost.json
// this example is just showing the usage of "Jobs_BookShipmentPickUp" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxJobResource created on azure
// for more information of creating DataBoxJobResource, please refer to the document of DataBoxJobResource
string subscriptionId = "YourSubscriptionId";
string resourceGroupName = "YourResourceGroupName";
string jobName = "TestJobName1";
ResourceIdentifier dataBoxJobResourceId = DataBoxJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
DataBoxJobResource dataBoxJob = client.GetDataBoxJobResource(dataBoxJobResourceId);

// invoke the operation
ShipmentPickUpContent content = new ShipmentPickUpContent(DateTimeOffset.Parse("2019-09-20T18:30:00Z"), DateTimeOffset.Parse("2019-09-22T18:30:00Z"), "Front desk");
DataBoxShipmentPickUpResult result = await dataBoxJob.BookShipmentPickUpAsync(content);

Console.WriteLine($"Succeeded: {result}");
