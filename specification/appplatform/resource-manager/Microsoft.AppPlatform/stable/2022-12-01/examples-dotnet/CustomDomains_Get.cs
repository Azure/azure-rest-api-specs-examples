using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/CustomDomains_Get.json
// this example is just showing the usage of "CustomDomains_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformAppResource created on azure
// for more information of creating AppPlatformAppResource, please refer to the document of AppPlatformAppResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string appName = "myapp";
ResourceIdentifier appPlatformAppResourceId = AppPlatformAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, appName);
AppPlatformAppResource appPlatformApp = client.GetAppPlatformAppResource(appPlatformAppResourceId);

// get the collection of this AppPlatformCustomDomainResource
AppPlatformCustomDomainCollection collection = appPlatformApp.GetAppPlatformCustomDomains();

// invoke the operation
string domainName = "mydomain.com";
NullableResponse<AppPlatformCustomDomainResource> response = await collection.GetIfExistsAsync(domainName);
AppPlatformCustomDomainResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AppPlatformCustomDomainData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
