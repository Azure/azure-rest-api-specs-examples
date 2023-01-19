using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridData;

// Generated from example definition: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_Cancel-POST-example-111.json
// this example is just showing the usage of "Jobs_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridDataJobResource created on azure
// for more information of creating HybridDataJobResource, please refer to the document of HybridDataJobResource
string subscriptionId = "6e0219f5-327a-4365-904f-05eed4227ad7";
string resourceGroupName = "ResourceGroupForSDKTest";
string dataManagerName = "TestAzureSDKOperations";
string dataServiceName = "DataTransformation";
string jobDefinitionName = "jobdeffromtestcode1";
string jobId = "6eca9b3d-5ffe-4b44-9607-1ba838371ff7";
ResourceIdentifier hybridDataJobResourceId = HybridDataJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataManagerName, dataServiceName, jobDefinitionName, jobId);
HybridDataJobResource hybridDataJob = client.GetHybridDataJobResource(hybridDataJobResourceId);

// invoke the operation
await hybridDataJob.CancelAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
