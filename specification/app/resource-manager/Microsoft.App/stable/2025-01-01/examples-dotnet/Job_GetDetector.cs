using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_GetDetector.json
// this example is just showing the usage of "Jobs_GetDetector" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppJobResource created on azure
// for more information of creating ContainerAppJobResource, please refer to the document of ContainerAppJobResource
string subscriptionId = "f07f3711-b45e-40fe-a941-4e6d93f851e6";
string resourceGroupName = "mikono-workerapp-test-rg";
string jobName = "mikonojob1";
ResourceIdentifier containerAppJobResourceId = ContainerAppJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
ContainerAppJobResource containerAppJob = client.GetContainerAppJobResource(containerAppJobResourceId);

// get the collection of this ContainerAppJobDetectorResource
ContainerAppJobDetectorCollection collection = containerAppJob.GetContainerAppJobDetectors();

// invoke the operation
string detectorName = "containerappjobnetworkIO";
NullableResponse<ContainerAppJobDetectorResource> response = await collection.GetIfExistsAsync(detectorName);
ContainerAppJobDetectorResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerAppDiagnosticData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
