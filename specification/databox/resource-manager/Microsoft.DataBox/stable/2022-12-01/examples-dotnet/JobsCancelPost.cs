using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataBox;
using Azure.ResourceManager.DataBox.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsCancelPost.json
// this example is just showing the usage of "Jobs_Cancel" operation, for the dependent resources, they will have to be created separately.

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
DataBoxJobCancellationReason cancellationReason = new DataBoxJobCancellationReason("CancelTest");
await dataBoxJob.CancelAsync(cancellationReason);

Console.WriteLine($"Succeeded");
