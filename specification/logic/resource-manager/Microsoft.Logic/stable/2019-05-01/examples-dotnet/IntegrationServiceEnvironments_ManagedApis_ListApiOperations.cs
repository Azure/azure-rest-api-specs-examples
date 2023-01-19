using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_ListApiOperations.json
// this example is just showing the usage of "IntegrationServiceEnvironmentManagedApiOperations_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationServiceEnvironmentManagedApiResource created on azure
// for more information of creating IntegrationServiceEnvironmentManagedApiResource, please refer to the document of IntegrationServiceEnvironmentManagedApiResource
string subscriptionId = "80d4fe69-c95b-4dd2-a938-9250f1c8ab03";
string resourceGroup = "testResourceGroup";
string integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
string apiName = "servicebus";
ResourceIdentifier integrationServiceEnvironmentManagedApiResourceId = IntegrationServiceEnvironmentManagedApiResource.CreateResourceIdentifier(subscriptionId, resourceGroup, integrationServiceEnvironmentName, apiName);
IntegrationServiceEnvironmentManagedApiResource integrationServiceEnvironmentManagedApi = client.GetIntegrationServiceEnvironmentManagedApiResource(integrationServiceEnvironmentManagedApiResourceId);

// invoke the operation and iterate over the result
await foreach (LogicApiOperationInfo item in integrationServiceEnvironmentManagedApi.GetIntegrationServiceEnvironmentManagedApiOperationsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
