using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ConfigureRemoteSupport.json
// this example is just showing the usage of "Clusters_ConfigureRemoteSupport" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterResource created on azure
// for more information of creating HciClusterResource, please refer to the document of HciClusterResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "mycluster";
ResourceIdentifier hciClusterResourceId = HciClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
HciClusterResource hciCluster = client.GetHciClusterResource(hciClusterResourceId);

// invoke the operation
RemoteSupportContent content = new RemoteSupportContent()
{
    Properties = new RemoteSupportContentProperties()
    {
        ExpireOn = DateTimeOffset.Parse("2020-01-01T17:18:19.1234567Z"),
        RemoteSupportType = RemoteSupportType.Enable,
    },
};
ArmOperation<HciClusterResource> lro = await hciCluster.ConfigureRemoteSupportAsync(WaitUntil.Completed, content);
HciClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
