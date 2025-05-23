using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetTimeseries.json
// this example is just showing the usage of "Reports_GetTimeSeries" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorExperimentResource created on azure
// for more information of creating FrontDoorExperimentResource, please refer to the document of FrontDoorExperimentResource
string subscriptionId = "subid";
string resourceGroupName = "MyResourceGroup";
string profileName = "MyProfile";
string experimentName = "MyExperiment";
ResourceIdentifier frontDoorExperimentResourceId = FrontDoorExperimentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, experimentName);
FrontDoorExperimentResource frontDoorExperiment = client.GetFrontDoorExperimentResource(frontDoorExperimentResourceId);

// invoke the operation
DateTimeOffset startOn = DateTimeOffset.Parse("2019-07-21T17:32:28Z");
DateTimeOffset endOn = DateTimeOffset.Parse("2019-09-21T17:32:28Z");
FrontDoorTimeSeriesAggregationInterval aggregationInterval = FrontDoorTimeSeriesAggregationInterval.Hourly;
FrontDoorTimeSeriesType timeSeriesType = FrontDoorTimeSeriesType.MeasurementCounts;
FrontDoorExperimentResourceGetTimeSeriesReportOptions options = new FrontDoorExperimentResourceGetTimeSeriesReportOptions(startOn, endOn, aggregationInterval, timeSeriesType);
FrontDoorTimeSeriesInfo result = await frontDoorExperiment.GetTimeSeriesReportAsync(options);

Console.WriteLine($"Succeeded: {result}");
