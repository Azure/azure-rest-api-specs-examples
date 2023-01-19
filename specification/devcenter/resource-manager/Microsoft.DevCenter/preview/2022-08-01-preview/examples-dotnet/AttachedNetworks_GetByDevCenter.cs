using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_GetByDevCenter.json
// this example is just showing the usage of "AttachedNetworks_GetByDevCenter" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AttachedNetworkConnectionResource created on azure
// for more information of creating AttachedNetworkConnectionResource, please refer to the document of AttachedNetworkConnectionResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
string attachedNetworkConnectionName = "network-uswest3";
ResourceIdentifier attachedNetworkConnectionResourceId = AttachedNetworkConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName, attachedNetworkConnectionName);
AttachedNetworkConnectionResource attachedNetworkConnection = client.GetAttachedNetworkConnectionResource(attachedNetworkConnectionResourceId);

// invoke the operation
AttachedNetworkConnectionResource result = await attachedNetworkConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AttachedNetworkConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
