using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_Put.json
// this example is just showing the usage of "IntegrationServiceEnvironmentManagedApis_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationServiceEnvironmentResource created on azure
// for more information of creating IntegrationServiceEnvironmentResource, please refer to the document of IntegrationServiceEnvironmentResource
string subscriptionId = "f34b22a3-2202-4fb1-b040-1332bd928c84";
string resourceGroup = "testResourceGroup";
string integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
ResourceIdentifier integrationServiceEnvironmentResourceId = IntegrationServiceEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroup, integrationServiceEnvironmentName);
IntegrationServiceEnvironmentResource integrationServiceEnvironment = client.GetIntegrationServiceEnvironmentResource(integrationServiceEnvironmentResourceId);

// get the collection of this IntegrationServiceEnvironmentManagedApiResource
IntegrationServiceEnvironmentManagedApiCollection collection = integrationServiceEnvironment.GetIntegrationServiceEnvironmentManagedApis();

// invoke the operation
string apiName = "servicebus";
IntegrationServiceEnvironmentManagedApiData data = new IntegrationServiceEnvironmentManagedApiData(new AzureLocation("brazilsouth"));
ArmOperation<IntegrationServiceEnvironmentManagedApiResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, apiName, data);
IntegrationServiceEnvironmentManagedApiResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IntegrationServiceEnvironmentManagedApiData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
