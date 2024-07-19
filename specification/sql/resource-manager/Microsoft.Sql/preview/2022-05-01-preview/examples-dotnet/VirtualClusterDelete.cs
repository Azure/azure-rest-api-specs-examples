using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/VirtualClusterDelete.json
// this example is just showing the usage of "VirtualClusters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualClusterResource created on azure
// for more information of creating VirtualClusterResource, please refer to the document of VirtualClusterResource
string subscriptionId = "20d7082a-0fc7-4468-82bd-542694d5042b";
string resourceGroupName = "testrg";
string virtualClusterName = "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2";
ResourceIdentifier virtualClusterResourceId = VirtualClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualClusterName);
VirtualClusterResource virtualCluster = client.GetVirtualClusterResource(virtualClusterResourceId);

// invoke the operation
await virtualCluster.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
