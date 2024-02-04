using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Purview;
using Azure.ResourceManager.Purview.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/preview/2023-05-01-preview/examples/IngestionPrivateEndpointConnections_UpdateStatus.json
// this example is just showing the usage of "IngestionPrivateEndpointConnections_UpdateStatus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PurviewAccountResource created on azure
// for more information of creating PurviewAccountResource, please refer to the document of PurviewAccountResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "SampleResourceGroup";
string accountName = "account1";
ResourceIdentifier purviewAccountResourceId = PurviewAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
PurviewAccountResource purviewAccount = client.GetPurviewAccountResource(purviewAccountResourceId);

// invoke the operation
PrivateEndpointConnectionStatusUpdateContent content = new PrivateEndpointConnectionStatusUpdateContent()
{
    PrivateEndpointId = "/subscriptions/12345678-1234-1234-12345678abc/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/privateEndpointConnections/privateEndpointConnection1",
    Status = "Approved",
};
PrivateEndpointConnectionStatusUpdateResult result = await purviewAccount.UpdateStatusIngestionPrivateEndpointConnectionAsync(content);

Console.WriteLine($"Succeeded: {result}");
