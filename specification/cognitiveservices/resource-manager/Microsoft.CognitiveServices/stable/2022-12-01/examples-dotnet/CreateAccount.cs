using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CognitiveServices;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/CreateAccount.json
// this example is just showing the usage of "Accounts_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CognitiveServicesAccountResource
CognitiveServicesAccountCollection collection = resourceGroupResource.GetCognitiveServicesAccounts();

// invoke the operation
string accountName = "testCreate1";
CognitiveServicesAccountData data = new CognitiveServicesAccountData(new AzureLocation("West US"))
{
    Kind = "Emotion",
    Sku = new CognitiveServicesSku("S0"),
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Properties = new CognitiveServicesAccountProperties()
    {
        Encryption = new ServiceAccountEncryptionProperties()
        {
            KeyVaultProperties = new CognitiveServicesKeyVaultProperties()
            {
                KeyName = "KeyName",
                KeyVersion = "891CF236-D241-4738-9462-D506AF493DFA",
                KeyVaultUri = new Uri("https://pltfrmscrts-use-pc-dev.vault.azure.net/"),
            },
            KeySource = ServiceAccountEncryptionKeySource.MicrosoftKeyVault,
        },
        UserOwnedStorage =
        {
        new ServiceAccountUserOwnedStorage()
        {
        ResourceId = new ResourceIdentifier("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
        }
        },
    },
};
ArmOperation<CognitiveServicesAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, data);
CognitiveServicesAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CognitiveServicesAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
