using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LoadTesting;
using Azure.ResourceManager.LoadTesting.Models;

// Generated from example definition: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_Get.json
// this example is just showing the usage of "Quotas_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LoadTestingQuotaResource created on azure
// for more information of creating LoadTestingQuotaResource, please refer to the document of LoadTestingQuotaResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
AzureLocation location = new AzureLocation("westus");
string quotaBucketName = "testQuotaBucket";
ResourceIdentifier loadTestingQuotaResourceId = LoadTestingQuotaResource.CreateResourceIdentifier(subscriptionId, location, quotaBucketName);
LoadTestingQuotaResource loadTestingQuota = client.GetLoadTestingQuotaResource(loadTestingQuotaResourceId);

// invoke the operation
LoadTestingQuotaResource result = await loadTestingQuota.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LoadTestingQuotaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
