using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_listPrivateIpAddresses.json
// this example is just showing the usage of "CloudVmClusters_ListPrivateIPAddresses" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudVmClusterResource created on azure
// for more information of creating CloudVmClusterResource, please refer to the document of CloudVmClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg000";
string cloudvmclustername = "cluster1";
ResourceIdentifier cloudVmClusterResourceId = CloudVmClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudvmclustername);
CloudVmClusterResource cloudVmCluster = client.GetCloudVmClusterResource(cloudVmClusterResourceId);

// invoke the operation and iterate over the result
PrivateIPAddressesContent content = new PrivateIPAddressesContent(new ResourceIdentifier("ocid1..aaaaaa"), new ResourceIdentifier("ocid1..aaaaa"));
await foreach (PrivateIPAddressResult item in cloudVmCluster.GetPrivateIPAddressesAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
