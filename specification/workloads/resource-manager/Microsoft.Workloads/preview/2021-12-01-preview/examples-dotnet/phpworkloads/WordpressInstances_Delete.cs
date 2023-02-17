using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/WordpressInstances_Delete.json
// this example is just showing the usage of "WordpressInstances_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WordPressInstanceResource created on azure
// for more information of creating WordPressInstanceResource, please refer to the document of WordPressInstanceResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
string resourceGroupName = "test-rg";
string phpWorkloadName = "wp39";
ResourceIdentifier wordPressInstanceResourceId = WordPressInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, phpWorkloadName);
WordPressInstanceResource wordPressInstanceResource = client.GetWordPressInstanceResource(wordPressInstanceResourceId);

// invoke the operation
await wordPressInstanceResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
