using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceNetworking.Models;
using Azure.ResourceManager.ServiceNetworking;

// Generated from example definition: 2025-01-01/AssociationDelete.json
// this example is just showing the usage of "Association_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficControllerAssociationResource created on azure
// for more information of creating TrafficControllerAssociationResource, please refer to the document of TrafficControllerAssociationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string trafficControllerName = "tc1";
string associationName = "as1";
ResourceIdentifier trafficControllerAssociationResourceId = TrafficControllerAssociationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, trafficControllerName, associationName);
TrafficControllerAssociationResource trafficControllerAssociation = client.GetTrafficControllerAssociationResource(trafficControllerAssociationResourceId);

// invoke the operation
await trafficControllerAssociation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
