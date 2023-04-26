using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/ProviderInstances_Get.json
// this example is just showing the usage of "ProviderInstances_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapProviderInstanceResource created on azure
// for more information of creating SapProviderInstanceResource, please refer to the document of SapProviderInstanceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "mySapMonitor";
string providerInstanceName = "myProviderInstance";
ResourceIdentifier sapProviderInstanceResourceId = SapProviderInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, providerInstanceName);
SapProviderInstanceResource sapProviderInstance = client.GetSapProviderInstanceResource(sapProviderInstanceResourceId);

// invoke the operation
SapProviderInstanceResource result = await sapProviderInstance.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapProviderInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
