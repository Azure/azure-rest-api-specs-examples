using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/WorkloadNetworks_DeleteDnsService.json
// this example is just showing the usage of "WorkloadNetworkDnsService_DeleteDnsService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadNetworkDnsServiceResource created on azure
// for more information of creating WorkloadNetworkDnsServiceResource, please refer to the document of WorkloadNetworkDnsServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string dnsServiceId = "dnsService1";
ResourceIdentifier workloadNetworkDnsServiceResourceId = WorkloadNetworkDnsServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, dnsServiceId);
WorkloadNetworkDnsServiceResource workloadNetworkDnsService = client.GetWorkloadNetworkDnsServiceResource(workloadNetworkDnsServiceResourceId);

// invoke the operation
await workloadNetworkDnsService.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
