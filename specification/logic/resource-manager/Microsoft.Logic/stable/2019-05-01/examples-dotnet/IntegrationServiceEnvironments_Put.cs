using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Put.json
// this example is just showing the usage of "IntegrationServiceEnvironments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "f34b22a3-2202-4fb1-b040-1332bd928c84";
string resourceGroup = "testResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroup);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IntegrationServiceEnvironmentResource
IntegrationServiceEnvironmentCollection collection = resourceGroupResource.GetIntegrationServiceEnvironments();

// invoke the operation
string integrationServiceEnvironmentName = "testIntegrationServiceEnvironment";
IntegrationServiceEnvironmentData data = new IntegrationServiceEnvironmentData(new AzureLocation("brazilsouth"))
{
    Properties = new IntegrationServiceEnvironmentProperties()
    {
        NetworkConfiguration = new IntegrationServiceNetworkConfiguration()
        {
            EndpointType = IntegrationServiceEnvironmentAccessEndpointType.Internal,
            Subnets =
            {
            new LogicResourceReference()
            {
            Id = new ResourceIdentifier("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s1"),
            },new LogicResourceReference()
            {
            Id = new ResourceIdentifier("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s2"),
            },new LogicResourceReference()
            {
            Id = new ResourceIdentifier("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s3"),
            },new LogicResourceReference()
            {
            Id = new ResourceIdentifier("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Network/virtualNetworks/testVNET/subnets/s4"),
            }
            },
        },
        EncryptionKeyReference = new IntegrationServiceEnvironmenEncryptionKeyReference()
        {
            KeyVault = new LogicResourceReference()
            {
                Id = new ResourceIdentifier("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault"),
            },
            KeyName = "testKeyName",
            KeyVersion = "13b261d30b984753869902d7f47f4d55",
        },
    },
    Sku = new IntegrationServiceEnvironmentSku()
    {
        Name = IntegrationServiceEnvironmentSkuName.Premium,
        Capacity = 2,
    },
};
ArmOperation<IntegrationServiceEnvironmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, integrationServiceEnvironmentName, data);
IntegrationServiceEnvironmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IntegrationServiceEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
