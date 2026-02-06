using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceGroups.Models;
using Azure.ResourceManager.ServiceGroups;

// Generated from example definition: specification/management/resource-manager/Microsoft.Management/ServiceGroups/preview/2024-02-01-preview/examples/ServiceGroup_Delete.json
// this example is just showing the usage of "ServiceGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceGroupResource created on azure
// for more information of creating ServiceGroupResource, please refer to the document of ServiceGroupResource
string serviceGroupName = "20000000-0001-0000-0000-000000000000";
ResourceIdentifier serviceGroupResourceId = ServiceGroupResource.CreateResourceIdentifier(serviceGroupName);
ServiceGroupResource serviceGroup = client.GetServiceGroupResource(serviceGroupResourceId);

// invoke the operation
await serviceGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
