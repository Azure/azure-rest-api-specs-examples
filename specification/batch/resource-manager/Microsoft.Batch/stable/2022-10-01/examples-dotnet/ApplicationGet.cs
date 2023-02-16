using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/ApplicationGet.json
// this example is just showing the usage of "Application_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchApplicationResource created on azure
// for more information of creating BatchApplicationResource, please refer to the document of BatchApplicationResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string applicationName = "app1";
ResourceIdentifier batchApplicationResourceId = BatchApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, applicationName);
BatchApplicationResource batchApplication = client.GetBatchApplicationResource(batchApplicationResourceId);

// invoke the operation
BatchApplicationResource result = await batchApplication.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
