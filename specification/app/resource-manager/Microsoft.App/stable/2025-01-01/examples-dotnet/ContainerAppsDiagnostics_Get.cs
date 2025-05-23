using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ContainerAppsDiagnostics_Get.json
// this example is just showing the usage of "ContainerAppsDiagnostics_GetDetector" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppDetectorResource created on azure
// for more information of creating ContainerAppDetectorResource, please refer to the document of ContainerAppDetectorResource
string subscriptionId = "f07f3711-b45e-40fe-a941-4e6d93f851e6";
string resourceGroupName = "mikono-workerapp-test-rg";
string containerAppName = "mikono-capp-stage1";
string detectorName = "cappcontainerappnetworkIO";
ResourceIdentifier containerAppDetectorResourceId = ContainerAppDetectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerAppName, detectorName);
ContainerAppDetectorResource containerAppDetector = client.GetContainerAppDetectorResource(containerAppDetectorResourceId);

// invoke the operation
ContainerAppDetectorResource result = await containerAppDetector.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppDiagnosticData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
