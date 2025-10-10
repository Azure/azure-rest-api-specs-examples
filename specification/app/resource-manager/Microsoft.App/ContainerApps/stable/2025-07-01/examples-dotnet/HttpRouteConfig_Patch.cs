using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/HttpRouteConfig_Patch.json
// this example is just showing the usage of "HttpRouteConfig_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppHttpRouteConfigResource created on azure
// for more information of creating ContainerAppHttpRouteConfigResource, please refer to the document of ContainerAppHttpRouteConfigResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "examplerg";
string environmentName = "testcontainerenv";
string httpRouteName = "httproutefriendlyname";
ResourceIdentifier containerAppHttpRouteConfigResourceId = ContainerAppHttpRouteConfigResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName, httpRouteName);
ContainerAppHttpRouteConfigResource containerAppHttpRouteConfig = client.GetContainerAppHttpRouteConfigResource(containerAppHttpRouteConfigResourceId);

// invoke the operation
ContainerAppHttpRouteConfigData data = new ContainerAppHttpRouteConfigData
{
    Properties = new ContainerAppHttpRouteConfigProperties
    {
        CustomDomains = {new ContainerAppCustomDomain("example.com")
        {
        BindingType = ContainerAppCustomDomainBindingType.SniEnabled,
        CertificateId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-1"),
        }},
        Rules = {new ContainerAppHttpRouteRule
        {
        Targets = {new ContainerAppHttpRouteTarget("capp-1")
        {
        Revision = "capp-1--0000001",
        }},
        Routes = {new ContainerAppHttpRoute
        {
        Match = new ContainerAppHttpRouteMatch
        {
        Path = "/v1",
        IsCaseSensitive = true,
        },
        ActionPrefixRewrite = "/v1/api",
        }},
        Description = "random-description",
        }},
    },
};
ContainerAppHttpRouteConfigResource result = await containerAppHttpRouteConfig.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppHttpRouteConfigData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
