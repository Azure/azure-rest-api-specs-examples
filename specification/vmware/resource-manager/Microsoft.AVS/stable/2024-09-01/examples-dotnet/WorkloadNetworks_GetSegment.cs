using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/WorkloadNetworks_GetSegment.json
// this example is just showing the usage of "WorkloadNetworkSegment_GetSegment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadNetworkSegmentResource created on azure
// for more information of creating WorkloadNetworkSegmentResource, please refer to the document of WorkloadNetworkSegmentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string segmentId = "segment1";
ResourceIdentifier workloadNetworkSegmentResourceId = WorkloadNetworkSegmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, segmentId);
WorkloadNetworkSegmentResource workloadNetworkSegment = client.GetWorkloadNetworkSegmentResource(workloadNetworkSegmentResourceId);

// invoke the operation
WorkloadNetworkSegmentResource result = await workloadNetworkSegment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkloadNetworkSegmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
