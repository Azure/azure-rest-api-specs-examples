using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Skus.json
// this example is just showing the usage of "IntegrationServiceEnvironmentSkus_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationServiceEnvironmentResource created on azure
// for more information of creating IntegrationServiceEnvironmentResource, please refer to the document of IntegrationServiceEnvironmentResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroup = "testResourceGroup";
string integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
ResourceIdentifier integrationServiceEnvironmentResourceId = IntegrationServiceEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroup, integrationServiceEnvironmentName);
IntegrationServiceEnvironmentResource integrationServiceEnvironment = client.GetIntegrationServiceEnvironmentResource(integrationServiceEnvironmentResourceId);

// invoke the operation and iterate over the result
await foreach (IntegrationServiceEnvironmentSkuDefinition item in integrationServiceEnvironment.GetIntegrationServiceEnvironmentSkusAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
