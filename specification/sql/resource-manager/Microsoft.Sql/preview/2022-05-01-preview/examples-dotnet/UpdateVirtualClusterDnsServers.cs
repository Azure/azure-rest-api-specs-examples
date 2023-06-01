using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/UpdateVirtualClusterDnsServers.json
// this example is just showing the usage of "VirtualClusters_UpdateDnsServers" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualClusterResource created on azure
// for more information of creating VirtualClusterResource, please refer to the document of VirtualClusterResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string virtualClusterName = "VirtualCluster2b9a846b-2e37-43ef-a8e9-f2c6d645c1d7";
ResourceIdentifier virtualClusterResourceId = VirtualClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualClusterName);
VirtualClusterResource virtualCluster = client.GetVirtualClusterResource(virtualClusterResourceId);

// invoke the operation
ArmOperation<ManagedInstanceUpdateDnsServersOperationData> lro = await virtualCluster.UpdateDnsServersAsync(WaitUntil.Completed);
ManagedInstanceUpdateDnsServersOperationData result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
