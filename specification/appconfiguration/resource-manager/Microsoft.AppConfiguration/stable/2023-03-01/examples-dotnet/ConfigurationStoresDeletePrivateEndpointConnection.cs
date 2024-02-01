using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppConfiguration;
using Azure.ResourceManager.AppConfiguration.Models;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresDeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppConfigurationPrivateEndpointConnectionResource created on azure
// for more information of creating AppConfigurationPrivateEndpointConnectionResource, please refer to the document of AppConfigurationPrivateEndpointConnectionResource
string subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
string resourceGroupName = "myResourceGroup";
string configStoreName = "contoso";
string privateEndpointConnectionName = "myConnection";
ResourceIdentifier appConfigurationPrivateEndpointConnectionResourceId = AppConfigurationPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configStoreName, privateEndpointConnectionName);
AppConfigurationPrivateEndpointConnectionResource appConfigurationPrivateEndpointConnection = client.GetAppConfigurationPrivateEndpointConnectionResource(appConfigurationPrivateEndpointConnectionResourceId);

// invoke the operation
await appConfigurationPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
