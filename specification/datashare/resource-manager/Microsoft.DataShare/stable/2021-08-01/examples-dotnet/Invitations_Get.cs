using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/Invitations_Get.json
// this example is just showing the usage of "Invitations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataShareResource created on azure
// for more information of creating DataShareResource, please refer to the document of DataShareResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareName = "Share1";
ResourceIdentifier dataShareResourceId = DataShareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName);
DataShareResource dataShare = client.GetDataShareResource(dataShareResourceId);

// get the collection of this DataShareInvitationResource
DataShareInvitationCollection collection = dataShare.GetDataShareInvitations();

// invoke the operation
string invitationName = "Invitation1";
bool result = await collection.ExistsAsync(invitationName);

Console.WriteLine($"Succeeded: {result}");
