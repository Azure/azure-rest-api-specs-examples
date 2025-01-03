using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthDataAIServices.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HealthDataAIServices;

// Generated from example definition: 2024-09-20/DeidServices_Create_MaximumSet_Gen.json
// this example is just showing the usage of "DeidService_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "F21BB31B-C214-42C0-ACF0-DACCA05D3011";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DeidServiceResource
DeidServiceCollection collection = resourceGroupResource.GetDeidServices();

// invoke the operation
string deidServiceName = "deidTest";
DeidServiceData data = new DeidServiceData(new AzureLocation("qwyhvdwcsjulggagdqxlmazcl"))
{
    Properties = new DeidServiceProperties()
    {
        PublicNetworkAccess = HealthDataAIServicesPublicNetworkAccess.Enabled,
    },
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        },
    },
    Tags =
    {
    },
};
ArmOperation<DeidServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, deidServiceName, data);
DeidServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeidServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
