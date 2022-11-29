using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudServiceUpdateDomain_Update.json
// this example is just showing the usage of "CloudServicesUpdateDomain_WalkUpdateDomain" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this CloudServiceResource created on azure
// for more information of creating CloudServiceResource, please refer to the document of CloudServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "ConstosoRG";
string cloudServiceName = "{cs-name}";
ResourceIdentifier cloudServiceResourceId = CloudServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudServiceName);
CloudServiceResource cloudService = client.GetCloudServiceResource(cloudServiceResourceId);

// invoke the operation
int updateDomain = 1;
await cloudService.WalkUpdateDomainAsync(WaitUntil.Completed, updateDomain);

Console.WriteLine($"Succeeded");
