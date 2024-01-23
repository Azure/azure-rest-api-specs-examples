using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;
using Azure.ResourceManager.Avs.Models;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/WorkloadNetworks_CreateDnsServices.json
// this example is just showing the usage of "WorkloadNetworks_CreateDnsService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudResource created on azure
// for more information of creating AvsPrivateCloudResource, please refer to the document of AvsPrivateCloudResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
ResourceIdentifier avsPrivateCloudResourceId = AvsPrivateCloudResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName);
AvsPrivateCloudResource avsPrivateCloud = client.GetAvsPrivateCloudResource(avsPrivateCloudResourceId);

// get the collection of this WorkloadNetworkDnsServiceResource
WorkloadNetworkDnsServiceCollection collection = avsPrivateCloud.GetWorkloadNetworkDnsServices();

// invoke the operation
string dnsServiceId = "dnsService1";
WorkloadNetworkDnsServiceData data = new WorkloadNetworkDnsServiceData()
{
    DisplayName = "dnsService1",
    DnsServiceIP = IPAddress.Parse("5.5.5.5"),
    DefaultDnsZone = "defaultDnsZone1",
    FqdnZones =
    {
    "fqdnZone1"
    },
    LogLevel = DnsServiceLogLevel.Info,
    Revision = 1,
};
ArmOperation<WorkloadNetworkDnsServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dnsServiceId, data);
WorkloadNetworkDnsServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkloadNetworkDnsServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
