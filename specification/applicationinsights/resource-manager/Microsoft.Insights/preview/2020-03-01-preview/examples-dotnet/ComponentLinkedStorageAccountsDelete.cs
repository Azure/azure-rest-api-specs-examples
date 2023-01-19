using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsDelete.json
// this example is just showing the usage of "ComponentLinkedStorageAccounts_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ComponentLinkedStorageAccountResource created on azure
// for more information of creating ComponentLinkedStorageAccountResource, please refer to the document of ComponentLinkedStorageAccountResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
string resourceGroupName = "someResourceGroupName";
string resourceName = "myComponent";
StorageType storageType = StorageType.ServiceProfiler;
ResourceIdentifier componentLinkedStorageAccountResourceId = ComponentLinkedStorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, storageType);
ComponentLinkedStorageAccountResource componentLinkedStorageAccount = client.GetComponentLinkedStorageAccountResource(componentLinkedStorageAccountResourceId);

// invoke the operation
await componentLinkedStorageAccount.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
