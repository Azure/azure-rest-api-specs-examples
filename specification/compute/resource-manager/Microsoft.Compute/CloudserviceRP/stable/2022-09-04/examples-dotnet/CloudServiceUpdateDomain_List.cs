using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceUpdateDomain_List.json
// this example is just showing the usage of "CloudServicesUpdateDomain_ListUpdateDomains" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudServiceResource created on azure
// for more information of creating CloudServiceResource, please refer to the document of CloudServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "ConstosoRG";
string cloudServiceName = "{cs-name}";
ResourceIdentifier cloudServiceResourceId = CloudServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudServiceName);
CloudServiceResource cloudService = client.GetCloudServiceResource(cloudServiceResourceId);

// invoke the operation and iterate over the result
await foreach (UpdateDomainIdentifier item in cloudService.GetUpdateDomainsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
