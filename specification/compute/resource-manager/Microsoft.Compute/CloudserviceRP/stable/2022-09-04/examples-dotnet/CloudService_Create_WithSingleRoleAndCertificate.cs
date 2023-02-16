using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudService_Create_WithSingleRoleAndCertificate.json
// this example is just showing the usage of "CloudServices_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "ConstosoRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CloudServiceResource
CloudServiceCollection collection = resourceGroupResource.GetCloudServices();

// invoke the operation
string cloudServiceName = "{cs-name}";
CloudServiceData data = new CloudServiceData(new AzureLocation("westus"))
{
    PackageUri = new Uri("{PackageUrl}"),
    Configuration = "{ServiceConfiguration}",
    UpgradeMode = CloudServiceUpgradeMode.Auto,
    Roles =
    {
    new CloudServiceRoleProfileProperties()
    {
    Name = "ContosoFrontend",
    Sku = new CloudServiceRoleSku()
    {
    Name = "Standard_D1_v2",
    Tier = "Standard",
    Capacity = 1,
    },
    }
    },
    OSSecrets =
    {
    new CloudServiceVaultSecretGroup()
    {
    SourceVaultId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.KeyVault/vaults/{keyvault-name}"),
    VaultCertificates =
    {
    new CloudServiceVaultCertificate()
    {
    CertificateUri = new Uri("https://{keyvault-name}.vault.azure.net:443/secrets/ContosoCertificate/{secret-id}"),
    }
    },
    }
    },
    NetworkProfile = new CloudServiceNetworkProfile()
    {
        LoadBalancerConfigurations =
        {
        new CloudServiceLoadBalancerConfiguration("contosolb",new LoadBalancerFrontendIPConfiguration[]
        {
        new LoadBalancerFrontendIPConfiguration("contosofe")
        {
        PublicIPAddressId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip"),
        }
        })
        },
    },
};
ArmOperation<CloudServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, cloudServiceName, data);
CloudServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
