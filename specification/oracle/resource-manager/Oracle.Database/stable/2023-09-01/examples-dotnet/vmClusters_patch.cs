using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_patch.json
// this example is just showing the usage of "CloudVmClusters_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
CloudVmClusterPatch patch = new CloudVmClusterPatch();
ArmOperation<CloudVmClusterResource> lro = await cloudVmCluster.UpdateAsync(WaitUntil.Completed, patch);
CloudVmClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudVmClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
