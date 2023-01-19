using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_Delete.json
// this example is just showing the usage of "AttachedNetworks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AttachedNetworkConnectionResource created on azure
// for more information of creating AttachedNetworkConnectionResource, please refer to the document of AttachedNetworkConnectionResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
string attachedNetworkConnectionName = "{attachedNetworkConnectionName}";
ResourceIdentifier attachedNetworkConnectionResourceId = AttachedNetworkConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName, attachedNetworkConnectionName);
AttachedNetworkConnectionResource attachedNetworkConnection = client.GetAttachedNetworkConnectionResource(attachedNetworkConnectionResourceId);

// invoke the operation
await attachedNetworkConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
