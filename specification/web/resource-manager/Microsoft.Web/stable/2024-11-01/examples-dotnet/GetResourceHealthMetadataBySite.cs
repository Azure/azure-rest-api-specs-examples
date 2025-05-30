using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetResourceHealthMetadataBySite.json
// this example is just showing the usage of "ResourceHealthMetadata_GetBySite" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteResourceHealthMetadataResource created on azure
// for more information of creating WebSiteResourceHealthMetadataResource, please refer to the document of WebSiteResourceHealthMetadataResource
string subscriptionId = "4adb32ad-8327-4cbb-b775-b84b4465bb38";
string resourceGroupName = "Default-Web-NorthCentralUS";
string name = "newsiteinnewASE-NCUS";
ResourceIdentifier webSiteResourceHealthMetadataResourceId = WebSiteResourceHealthMetadataResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
WebSiteResourceHealthMetadataResource webSiteResourceHealthMetadata = client.GetWebSiteResourceHealthMetadataResource(webSiteResourceHealthMetadataResourceId);

// invoke the operation
WebSiteResourceHealthMetadataResource result = await webSiteResourceHealthMetadata.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceHealthMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
