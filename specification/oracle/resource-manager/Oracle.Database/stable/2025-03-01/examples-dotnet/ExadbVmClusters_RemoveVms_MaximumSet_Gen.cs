using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-03-01/ExadbVmClusters_RemoveVms_MaximumSet_Gen.json
// this example is just showing the usage of "ExadbVmClusters_RemoveVms" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExadbVmClusterResource created on azure
// for more information of creating ExadbVmClusterResource, please refer to the document of ExadbVmClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgopenapi";
string exadbVmClusterName = "vmClusterName";
ResourceIdentifier exadbVmClusterResourceId = ExadbVmClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, exadbVmClusterName);
ExadbVmClusterResource exadbVmCluster = client.GetExadbVmClusterResource(exadbVmClusterResourceId);

// invoke the operation
RemoveVirtualMachineFromExadbVmClusterDetails details = new RemoveVirtualMachineFromExadbVmClusterDetails(new DBNodeDetails[]
{
new DBNodeDetails(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Oracle.Database/exadbVmClusters/vmCluster/dbNodes/dbNodeName"))
});
ArmOperation<ExadbVmClusterResource> lro = await exadbVmCluster.RemoveVmsAsync(WaitUntil.Completed, details);
ExadbVmClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExadbVmClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
