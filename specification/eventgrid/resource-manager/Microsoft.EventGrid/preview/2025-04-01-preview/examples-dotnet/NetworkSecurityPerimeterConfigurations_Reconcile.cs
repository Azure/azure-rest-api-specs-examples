using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/NetworkSecurityPerimeterConfigurations_Reconcile.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Reconcile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DomainNetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating DomainNetworkSecurityPerimeterConfigurationResource, please refer to the document of DomainNetworkSecurityPerimeterConfigurationResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string resourceName = "exampleResourceName";
string perimeterGuid = "8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter";
string associationName = "someAssociation";
ResourceIdentifier domainNetworkSecurityPerimeterConfigurationResourceId = DomainNetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, perimeterGuid, associationName);
DomainNetworkSecurityPerimeterConfigurationResource domainNetworkSecurityPerimeterConfiguration = client.GetDomainNetworkSecurityPerimeterConfigurationResource(domainNetworkSecurityPerimeterConfigurationResourceId);

// invoke the operation
ArmOperation<DomainNetworkSecurityPerimeterConfigurationResource> lro = await domainNetworkSecurityPerimeterConfiguration.ReconcileAsync(WaitUntil.Completed);
DomainNetworkSecurityPerimeterConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkSecurityPerimeterConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
