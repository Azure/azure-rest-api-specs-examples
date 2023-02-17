using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridData;

// Generated from example definition: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/Jobs_Get-GET-example-101.json
// this example is just showing the usage of "Jobs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridDataJobDefinitionResource created on azure
// for more information of creating HybridDataJobDefinitionResource, please refer to the document of HybridDataJobDefinitionResource
string subscriptionId = "6e0219f5-327a-4365-904f-05eed4227ad7";
string resourceGroupName = "ResourceGroupForSDKTest";
string dataManagerName = "TestAzureSDKOperations";
string dataServiceName = "DataTransformation";
string jobDefinitionName = "jobdeffromtestcode1";
ResourceIdentifier hybridDataJobDefinitionResourceId = HybridDataJobDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataManagerName, dataServiceName, jobDefinitionName);
HybridDataJobDefinitionResource hybridDataJobDefinition = client.GetHybridDataJobDefinitionResource(hybridDataJobDefinitionResourceId);

// get the collection of this HybridDataJobResource
HybridDataJobCollection collection = hybridDataJobDefinition.GetHybridDataJobs();

// invoke the operation
string jobId = "99ef93fe-36be-43e4-bebf-de6746730601";
bool result = await collection.ExistsAsync(jobId);

Console.WriteLine($"Succeeded: {result}");
