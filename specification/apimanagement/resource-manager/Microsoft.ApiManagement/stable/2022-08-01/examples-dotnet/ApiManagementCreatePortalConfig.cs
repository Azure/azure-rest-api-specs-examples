using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreatePortalConfig.json
// this example is just showing the usage of "PortalConfig_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementServiceResource created on azure
// for more information of creating ApiManagementServiceResource, please refer to the document of ApiManagementServiceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementServiceResourceId = ApiManagementServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementServiceResource apiManagementService = client.GetApiManagementServiceResource(apiManagementServiceResourceId);

// get the collection of this PortalConfigContractResource
PortalConfigContractCollection collection = apiManagementService.GetPortalConfigContracts();

// invoke the operation
string portalConfigId = "default";
ETag ifMatch = new ETag("*");
PortalConfigContractData data = new PortalConfigContractData()
{
    EnableBasicAuth = true,
    Require = false,
    SignupTermsOfService = new PortalConfigTermsOfServiceProperties()
    {
        Text = "I agree to the service terms and conditions.",
        RequireConsent = false,
    },
    Delegation = new PortalConfigDelegationProperties()
    {
        DelegateRegistration = false,
        DelegateSubscription = false,
        DelegationUri = null,
        ValidationKey = null,
    },
    CorsAllowedOrigins =
    {
    "https://contoso.com"
    },
    Csp = new PortalConfigCspProperties()
    {
        Mode = PortalSettingsCspMode.ReportOnly,
        ReportUri =
        {
        new Uri("https://report.contoso.com")
        },
        AllowedSources =
        {
        "*.contoso.com"
        },
    },
};
ArmOperation<PortalConfigContractResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, portalConfigId, ifMatch, data);
PortalConfigContractResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PortalConfigContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
