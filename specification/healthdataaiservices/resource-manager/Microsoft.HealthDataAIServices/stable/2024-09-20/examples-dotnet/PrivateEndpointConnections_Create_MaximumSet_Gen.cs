using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthDataAIServices.Models;
using Azure.ResourceManager.HealthDataAIServices;

// Generated from example definition: 2024-09-20/PrivateEndpointConnections_Create_MaximumSet_Gen.json
// this example is just showing the usage of "PrivateEndpointConnectionResource_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthDataAIServicesPrivateEndpointConnectionResource created on azure
// for more information of creating HealthDataAIServicesPrivateEndpointConnectionResource, please refer to the document of HealthDataAIServicesPrivateEndpointConnectionResource
string subscriptionId = "F21BB31B-C214-42C0-ACF0-DACCA05D3011";
string resourceGroupName = "rgopenapi";
string deidServiceName = "deidTest";
string privateEndpointConnectionName = "kgwgrrpabvrsrrvpcgcnfmyfgyrl";
ResourceIdentifier healthDataAIServicesPrivateEndpointConnectionResourceId = HealthDataAIServicesPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deidServiceName, privateEndpointConnectionName);
HealthDataAIServicesPrivateEndpointConnectionResource healthDataAIServicesPrivateEndpointConnectionResource = client.GetHealthDataAIServicesPrivateEndpointConnectionResource(healthDataAIServicesPrivateEndpointConnectionResourceId);

// invoke the operation
HealthDataAIServicesPrivateEndpointConnectionResourceData data = new HealthDataAIServicesPrivateEndpointConnectionResourceData()
{
    Properties = new PrivateEndpointConnectionProperties(new HealthDataAIServicesPrivateLinkServiceConnectionState()
    {
        Status = HealthDataAIServicesPrivateEndpointServiceConnectionStatus.Pending,
        Description = "xr",
        ActionsRequired = "ulb",
    }),
};
ArmOperation<HealthDataAIServicesPrivateEndpointConnectionResource> lro = await healthDataAIServicesPrivateEndpointConnectionResource.UpdateAsync(WaitUntil.Completed, data);
HealthDataAIServicesPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthDataAIServicesPrivateEndpointConnectionResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
