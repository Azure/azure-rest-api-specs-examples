using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthcareApis.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HealthcareApis;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/ServiceCreateInDataSovereignRegionWithCmkEnabled.json
// this example is just showing the usage of "Services_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HealthcareApisServiceResource
HealthcareApisServiceCollection collection = resourceGroupResource.GetHealthcareApisServices();

// invoke the operation
string resourceName = "service1";
HealthcareApisServiceData data = new HealthcareApisServiceData(new AzureLocation("Southeast Asia"), HealthcareApisKind.FhirR4)
{
    Properties = new HealthcareApisServiceProperties()
    {
        AccessPolicies =
        {
        new HealthcareApisServiceAccessPolicyEntry("c487e7d1-3210-41a3-8ccc-e9372b78da47"),new HealthcareApisServiceAccessPolicyEntry("5b307da8-43d4-492b-8b66-b0294ade872f")
        },
        CosmosDbConfiguration = new HealthcareApisServiceCosmosDbConfiguration()
        {
            OfferThroughput = 1000,
            KeyVaultKeyUri = new Uri("https://my-vault.vault.azure.net/keys/my-key"),
            CrossTenantCmkApplicationId = Guid.Parse("de3fbeef-8c3a-428e-8b9f-4d229c8a85f4"),
        },
        AuthenticationConfiguration = new HealthcareApisServiceAuthenticationConfiguration()
        {
            Authority = "https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc",
            Audience = "https://azurehealthcareapis.com",
            IsSmartProxyEnabled = true,
        },
        CorsConfiguration = new HealthcareApisServiceCorsConfiguration()
        {
            Origins =
            {
            "*"
            },
            Headers =
            {
            "*"
            },
            Methods =
            {
            "DELETE","GET","OPTIONS","PATCH","POST","PUT"
            },
            MaxAge = 1440,
            AllowCredentials = false,
        },
        ExportStorageAccountName = "existingStorageAccount",
        PrivateEndpointConnections =
        {
        },
        PublicNetworkAccess = HealthcareApisPublicNetworkAccess.Disabled,
    },
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Tags =
    {
    },
};
ArmOperation<HealthcareApisServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
HealthcareApisServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthcareApisServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
