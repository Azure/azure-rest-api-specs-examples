using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ProviderHub;
using Azure.ResourceManager.ProviderHub.Models;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_CreateOrUpdate.json
// this example is just showing the usage of "ResourceTypeRegistrations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProviderRegistrationResource created on azure
// for more information of creating ProviderRegistrationResource, please refer to the document of ProviderRegistrationResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
ResourceIdentifier providerRegistrationResourceId = ProviderRegistrationResource.CreateResourceIdentifier(subscriptionId, providerNamespace);
ProviderRegistrationResource providerRegistration = client.GetProviderRegistrationResource(providerRegistrationResourceId);

// get the collection of this ResourceTypeRegistrationResource
ResourceTypeRegistrationCollection collection = providerRegistration.GetResourceTypeRegistrations();

// invoke the operation
string resourceType = "employees";
ResourceTypeRegistrationData data = new ResourceTypeRegistrationData()
{
    Properties = new ResourceTypeRegistrationProperties()
    {
        RoutingType = ResourceRoutingType.Default,
        Regionality = ResourceTypeRegistrationRegionality.Regional,
        Endpoints =
        {
        new ResourceTypeEndpoint()
        {
        ApiVersions =
        {
        "2020-06-01-preview"
        },
        Locations =
        {
        new AzureLocation("West US"),new AzureLocation("East US"),new AzureLocation("North Europe")
        },
        RequiredFeatures =
        {
        "<feature flag>"
        },
        }
        },
        SwaggerSpecifications =
        {
        new SwaggerSpecification()
        {
        ApiVersions =
        {
        "2020-06-01-preview"
        },
        SwaggerSpecFolderUri = new Uri("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
        }
        },
    },
};
ArmOperation<ResourceTypeRegistrationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceType, data);
ResourceTypeRegistrationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceTypeRegistrationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
