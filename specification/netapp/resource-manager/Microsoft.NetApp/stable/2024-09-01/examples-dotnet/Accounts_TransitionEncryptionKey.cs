using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/Accounts_TransitionEncryptionKey.json
// this example is just showing the usage of "Accounts_TransitionToCmk" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppAccountResource created on azure
// for more information of creating NetAppAccountResource, please refer to the document of NetAppAccountResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
ResourceIdentifier netAppAccountResourceId = NetAppAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
NetAppAccountResource netAppAccount = client.GetNetAppAccountResource(netAppAccountResourceId);

// invoke the operation
NetAppEncryptionTransitionContent content = new NetAppEncryptionTransitionContent(new ResourceIdentifier("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.Network/virtualNetworks/vnet1"), new ResourceIdentifier("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.Network/privateEndpoints/privip1"));
await netAppAccount.TransitionToCmkAsync(WaitUntil.Completed, content: content);

Console.WriteLine("Succeeded");
