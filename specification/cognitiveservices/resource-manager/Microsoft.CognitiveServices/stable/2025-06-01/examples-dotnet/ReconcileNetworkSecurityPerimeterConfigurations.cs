using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/ReconcileNetworkSecurityPerimeterConfigurations.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Reconcile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating NetworkSecurityPerimeterConfigurationResource, please refer to the document of NetworkSecurityPerimeterConfigurationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string nspConfigurationName = "NSPConfigurationName";
ResourceIdentifier networkSecurityPerimeterConfigurationResourceId = NetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, nspConfigurationName);
NetworkSecurityPerimeterConfigurationResource networkSecurityPerimeterConfiguration = client.GetNetworkSecurityPerimeterConfigurationResource(networkSecurityPerimeterConfigurationResourceId);

// invoke the operation
ArmOperation<NetworkSecurityPerimeterConfigurationResource> lro = await networkSecurityPerimeterConfiguration.ReconcileAsync(WaitUntil.Completed);
NetworkSecurityPerimeterConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkSecurityPerimeterConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
