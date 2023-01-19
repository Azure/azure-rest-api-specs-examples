using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/Invitations_Delete.json
// this example is just showing the usage of "Invitations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataShareInvitationResource created on azure
// for more information of creating DataShareInvitationResource, please refer to the document of DataShareInvitationResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareName = "Share1";
string invitationName = "Invitation1";
ResourceIdentifier dataShareInvitationResourceId = DataShareInvitationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName, invitationName);
DataShareInvitationResource dataShareInvitation = client.GetDataShareInvitationResource(dataShareInvitationResourceId);

// invoke the operation
await dataShareInvitation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
