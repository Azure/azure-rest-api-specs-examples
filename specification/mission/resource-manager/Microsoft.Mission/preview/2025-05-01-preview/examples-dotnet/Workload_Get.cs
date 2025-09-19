using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/Workload_Get.json
// this example is just showing the usage of "WorkloadResource_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveWorkloadResource created on azure
// for more information of creating VirtualEnclaveWorkloadResource, please refer to the document of VirtualEnclaveWorkloadResource
string subscriptionId = "CA1CB369-DD26-4DB2-9D43-9AFEF0F22093";
string resourceGroupName = "rgopenapi";
string virtualEnclaveName = "TestMyEnclave";
string workloadName = "TestMyWorkload";
ResourceIdentifier virtualEnclaveWorkloadResourceId = VirtualEnclaveWorkloadResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualEnclaveName, workloadName);
VirtualEnclaveWorkloadResource virtualEnclaveWorkload = client.GetVirtualEnclaveWorkloadResource(virtualEnclaveWorkloadResourceId);

// invoke the operation
VirtualEnclaveWorkloadResource result = await virtualEnclaveWorkload.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveWorkloadData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
