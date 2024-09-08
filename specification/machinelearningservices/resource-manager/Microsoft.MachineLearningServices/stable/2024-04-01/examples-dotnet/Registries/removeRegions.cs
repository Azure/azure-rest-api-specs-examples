using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registries/removeRegions.json
// this example is just showing the usage of "Registries_RemoveRegions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryResource created on azure
// for more information of creating MachineLearningRegistryResource, please refer to the document of MachineLearningRegistryResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "string";
ResourceIdentifier machineLearningRegistryResourceId = MachineLearningRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName);
MachineLearningRegistryResource machineLearningRegistry = client.GetMachineLearningRegistryResource(machineLearningRegistryResourceId);

// invoke the operation
MachineLearningRegistryData data = new MachineLearningRegistryData(new AzureLocation("string"))
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("string")] = new UserAssignedIdentity(),
        },
    },
    Kind = "string",
    Sku = new MachineLearningSku("string")
    {
        Tier = MachineLearningSkuTier.Free,
        Size = "string",
        Family = "string",
        Capacity = 1,
    },
    DiscoveryUri = new Uri("string"),
    IntellectualPropertyPublisher = "string",
    ManagedResourceId = new ResourceIdentifier("string"),
    MlFlowRegistryUri = new Uri("string"),
    RegistryPrivateEndpointConnections =
    {
    new RegistryPrivateEndpointConnection()
    {
    Id = new ResourceIdentifier("string"),
    Location = new AzureLocation("string"),
    GroupIds =
    {
    "string"
    },
    PrivateEndpoint = new RegistryPrivateEndpoint()
    {
    SubnetArmId = new ResourceIdentifier("string"),
    },
    RegistryPrivateLinkServiceConnectionState = new RegistryPrivateLinkServiceConnectionState()
    {
    ActionsRequired = "string",
    Description = "string",
    Status = EndpointServiceConnectionStatus.Approved,
    },
    ProvisioningState = "string",
    }
    },
    PublicNetworkAccess = "string",
    RegionDetails =
    {
    new RegistryRegionArmDetails()
    {
    AcrDetails =
    {
    new RegistryAcrDetails()
    {
    SystemCreatedAcrAccount = new SystemCreatedAcrAccount()
    {
    AcrAccountName = "string",
    AcrAccountSku = "string",
    ArmResourceId = new ResourceIdentifier("string"),
    },
    ArmResourceId = new ResourceIdentifier("string"),
    }
    },
    Location = new AzureLocation("string"),
    StorageAccountDetails =
    {
    new StorageAccountDetails()
    {
    SystemCreatedStorageAccount = new SystemCreatedStorageAccount()
    {
    AllowBlobPublicAccess = false,
    ArmResourceId = new ResourceIdentifier("string"),
    StorageAccountHnsEnabled = false,
    StorageAccountName = "string",
    StorageAccountType = "string",
    },
    ArmResourceId = new ResourceIdentifier("string"),
    }
    },
    }
    },
    Tags =
    {
    },
};
ArmOperation<MachineLearningRegistryResource> lro = await machineLearningRegistry.RemoveRegionsAsync(WaitUntil.Completed, data);
MachineLearningRegistryResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningRegistryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
