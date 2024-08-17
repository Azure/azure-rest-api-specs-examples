using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevOpsInfrastructure.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DevOpsInfrastructure;

// Generated from example definition: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/GetPool.json
// this example is just showing the usage of "Pools_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevOpsPoolResource created on azure
// for more information of creating DevOpsPoolResource, please refer to the document of DevOpsPoolResource
string subscriptionId = "a2e95d27-c161-4b61-bda4-11512c14c2c2";
string resourceGroupName = "rg";
string poolName = "pool";
ResourceIdentifier devOpsPoolResourceId = DevOpsPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, poolName);
DevOpsPoolResource devOpsPool = client.GetDevOpsPoolResource(devOpsPoolResourceId);

// invoke the operation
DevOpsPoolResource result = await devOpsPool.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevOpsPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
