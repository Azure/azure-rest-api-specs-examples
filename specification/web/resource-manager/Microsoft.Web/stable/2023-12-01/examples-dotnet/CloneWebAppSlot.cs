using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/CloneWebAppSlot.json
// this example is just showing the usage of "WebApps_CreateOrUpdateSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteResource created on azure
// for more information of creating WebSiteResource, please refer to the document of WebSiteResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "sitef6141";
ResourceIdentifier webSiteResourceId = WebSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
WebSiteResource webSite = client.GetWebSiteResource(webSiteResourceId);

// get the collection of this WebSiteSlotResource
WebSiteSlotCollection collection = webSite.GetWebSiteSlots();

// invoke the operation
string slot = "staging";
WebSiteData data = new WebSiteData(new AzureLocation("East US"))
{
    CloningInfo = new CloningInfo(new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/sites/srcsiteg478/slot/qa"))
    {
        CanOverwrite = false,
        CloneCustomHostNames = true,
        CloneSourceControl = true,
        SourceWebAppLocation = new AzureLocation("West Europe"),
        HostingEnvironment = "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg456/providers/Microsoft.Web/hostingenvironments/aseforsites",
        AppSettingsOverrides =
        {
        ["Setting1"] = "NewValue1",
        ["Setting3"] = "NewValue5",
        },
        ConfigureLoadBalancing = false,
    },
    Kind = "app",
};
ArmOperation<WebSiteSlotResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, slot, data);
WebSiteSlotResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebSiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
