using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ApiPortalCustomDomains_CreateOrUpdate.json
// this example is just showing the usage of "ApiPortalCustomDomains_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformApiPortalCustomDomainResource created on azure
// for more information of creating AppPlatformApiPortalCustomDomainResource, please refer to the document of AppPlatformApiPortalCustomDomainResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string apiPortalName = "default";
string domainName = "myDomainName";
ResourceIdentifier appPlatformApiPortalCustomDomainResourceId = AppPlatformApiPortalCustomDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiPortalName, domainName);
AppPlatformApiPortalCustomDomainResource appPlatformApiPortalCustomDomain = client.GetAppPlatformApiPortalCustomDomainResource(appPlatformApiPortalCustomDomainResourceId);

// invoke the operation
AppPlatformApiPortalCustomDomainData data = new AppPlatformApiPortalCustomDomainData
{
    ApiPortalCustomDomainThumbprint = "*",
};
ArmOperation<AppPlatformApiPortalCustomDomainResource> lro = await appPlatformApiPortalCustomDomain.UpdateAsync(WaitUntil.Completed, data);
AppPlatformApiPortalCustomDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformApiPortalCustomDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
