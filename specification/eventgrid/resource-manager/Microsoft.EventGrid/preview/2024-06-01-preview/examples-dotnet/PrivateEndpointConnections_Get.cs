using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/PrivateEndpointConnections_Get.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerNamespaceResource created on azure
// for more information of creating PartnerNamespaceResource, please refer to the document of PartnerNamespaceResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string parentName = "exampletopic1";
ResourceIdentifier partnerNamespaceResourceId = PartnerNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentName);
PartnerNamespaceResource partnerNamespace = client.GetPartnerNamespaceResource(partnerNamespaceResourceId);

// get the collection of this EventGridPartnerNamespacePrivateEndpointConnectionResource
EventGridPartnerNamespacePrivateEndpointConnectionCollection collection = partnerNamespace.GetEventGridPartnerNamespacePrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "BMTPE5.8A30D251-4C61-489D-A1AA-B37C4A329B8B";
NullableResponse<EventGridPartnerNamespacePrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
EventGridPartnerNamespacePrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EventGridPrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
