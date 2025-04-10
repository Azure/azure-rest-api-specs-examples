using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: 2024-09-01/PrivateEndpointConnectionProxy_Create.json
// this example is just showing the usage of "PrivateEndpointConnectionProxy_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationPrivateEndpointConnectionProxyResource created on azure
// for more information of creating DataReplicationPrivateEndpointConnectionProxyResource, please refer to the document of DataReplicationPrivateEndpointConnectionProxyResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgswagger_2024-09-01";
string vaultName = "4";
string privateEndpointConnectionProxyName = "d";
ResourceIdentifier dataReplicationPrivateEndpointConnectionProxyResourceId = DataReplicationPrivateEndpointConnectionProxyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, privateEndpointConnectionProxyName);
DataReplicationPrivateEndpointConnectionProxyResource dataReplicationPrivateEndpointConnectionProxy = client.GetDataReplicationPrivateEndpointConnectionProxyResource(dataReplicationPrivateEndpointConnectionProxyResourceId);

// invoke the operation
DataReplicationPrivateEndpointConnectionProxyData data = new DataReplicationPrivateEndpointConnectionProxyData
{
    Properties = new DataReplicationPrivateEndpointConnectionProxyProperties
    {
        RemotePrivateEndpoint = new RemotePrivateEndpoint("yipalno")
        {
            PrivateLinkServiceConnections = {new DataReplicationPrivateLinkServiceConnection
            {
            Name = "jqwntlzfsksl",
            GroupIds = {"hvejynjktikteipnioyeja"},
            RequestMessage = "bukgzpkvcvfbmcdmpcbiigbvugicqa",
            }},
            ManualPrivateLinkServiceConnections = {new DataReplicationPrivateLinkServiceConnection
            {
            Name = "jqwntlzfsksl",
            GroupIds = {"hvejynjktikteipnioyeja"},
            RequestMessage = "bukgzpkvcvfbmcdmpcbiigbvugicqa",
            }},
            PrivateLinkServiceProxies = {new DataReplicationPrivateLinkServiceProxy
            {
            Id = "nzqxevuyqeedrqnkbnlcyrrrbzxvl",
            RemotePrivateLinkServiceConnectionState = new DataReplicationPrivateLinkServiceConnectionState
            {
            Status = DataReplicationPrivateEndpointConnectionStatus.Approved,
            Description = "y",
            ActionsRequired = "afwbq",
            },
            RemotePrivateEndpointConnectionId = new ResourceIdentifier("ocunsgawjsqohkrcyxiv"),
            GroupConnectivityInformation = {new GroupConnectivityInformation
            {
            GroupId = "per",
            MemberName = "ybptuypgdqoxkuwqx",
            CustomerVisibleFqdns = {"vedcg"},
            InternalFqdn = "maqavwhxwzzhbzjbryyquvitmup",
            RedirectMapId = "pezncxcq",
            PrivateLinkServiceArmRegion = "rerkqqxinteevmlbrdkktaqhcch",
            }},
            }},
            ConnectionDetails = {new RemotePrivateEndpointConnectionDetails
            {
            Id = "lenqkogzkes",
            PrivateIPAddress = "cyiacdzzyqmxjpijjbwgasegehtqe",
            LinkIdentifier = "ravfufhkdowufd",
            GroupId = "pjrlygpadir",
            MemberName = "ybuysjrlfupewxe",
            }},
        },
    },
    ETag = new ETag("hruibxrezstxroxrxweh"),
};
ArmOperation<DataReplicationPrivateEndpointConnectionProxyResource> lro = await dataReplicationPrivateEndpointConnectionProxy.UpdateAsync(WaitUntil.Completed, data);
DataReplicationPrivateEndpointConnectionProxyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataReplicationPrivateEndpointConnectionProxyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
