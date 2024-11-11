using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ApproveRejectPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnection_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryPrivateEndpointConnectionResource created on azure
// for more information of creating DataFactoryPrivateEndpointConnectionResource, please refer to the document of DataFactoryPrivateEndpointConnectionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string privateEndpointConnectionName = "connection";
ResourceIdentifier dataFactoryPrivateEndpointConnectionResourceId = DataFactoryPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, privateEndpointConnectionName);
DataFactoryPrivateEndpointConnectionResource dataFactoryPrivateEndpointConnection = client.GetDataFactoryPrivateEndpointConnectionResource(dataFactoryPrivateEndpointConnectionResourceId);

// invoke the operation
DataFactoryPrivateEndpointConnectionCreateOrUpdateContent content = new DataFactoryPrivateEndpointConnectionCreateOrUpdateContent()
{
    Properties = new PrivateLinkConnectionApprovalRequest()
    {
        PrivateLinkServiceConnectionState = new PrivateLinkConnectionState()
        {
            Status = "Approved",
            Description = "Approved by admin.",
            ActionsRequired = "",
        },
        PrivateEndpointId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
    },
};
ArmOperation<DataFactoryPrivateEndpointConnectionResource> lro = await dataFactoryPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, content);
DataFactoryPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
