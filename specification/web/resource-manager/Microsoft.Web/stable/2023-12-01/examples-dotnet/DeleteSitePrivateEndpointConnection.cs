using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/DeleteSitePrivateEndpointConnection.json
// this example is just showing the usage of "StaticSites_DeletePrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSitePrivateEndpointConnectionResource created on azure
// for more information of creating StaticSitePrivateEndpointConnectionResource, please refer to the document of StaticSitePrivateEndpointConnectionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testSite";
string privateEndpointConnectionName = "connection";
ResourceIdentifier staticSitePrivateEndpointConnectionResourceId = StaticSitePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, privateEndpointConnectionName);
StaticSitePrivateEndpointConnectionResource staticSitePrivateEndpointConnection = client.GetStaticSitePrivateEndpointConnectionResource(staticSitePrivateEndpointConnectionResourceId);

// invoke the operation
ArmOperation<BinaryData> lro = await staticSitePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);
BinaryData result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
