using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevOpsInfrastructure.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DevOpsInfrastructure;

// Generated from example definition: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/ResourceDetails_ListByPool.json
// this example is just showing the usage of "ResourceDetails_ListByPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevOpsPoolResource created on azure
// for more information of creating DevOpsPoolResource, please refer to the document of DevOpsPoolResource
string subscriptionId = "a2e95d27-c161-4b61-bda4-11512c14c2c2";
string resourceGroupName = "my-resource-group";
string poolName = "my-dev-ops-pool";
ResourceIdentifier devOpsPoolResourceId = DevOpsPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, poolName);
DevOpsPoolResource devOpsPool = client.GetDevOpsPoolResource(devOpsPoolResourceId);

// invoke the operation and iterate over the result
await foreach (DevOpsResourceDetails item in devOpsPool.GetResourceDetailsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
