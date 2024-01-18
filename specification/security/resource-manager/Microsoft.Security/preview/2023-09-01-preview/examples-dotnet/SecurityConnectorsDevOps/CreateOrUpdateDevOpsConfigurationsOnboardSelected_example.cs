using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/CreateOrUpdateDevOpsConfigurationsOnboardSelected_example.json
// this example is just showing the usage of "DevOpsConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevOpsConfigurationResource created on azure
// for more information of creating DevOpsConfigurationResource, please refer to the document of DevOpsConfigurationResource
string subscriptionId = "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
string resourceGroupName = "myRg";
string securityConnectorName = "mySecurityConnectorName";
ResourceIdentifier devOpsConfigurationResourceId = DevOpsConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, securityConnectorName);
DevOpsConfigurationResource devOpsConfiguration = client.GetDevOpsConfigurationResource(devOpsConfigurationResourceId);

// invoke the operation
DevOpsConfigurationData data = new DevOpsConfigurationData()
{
    Properties = new DevOpsConfigurationProperties()
    {
        AuthorizationCode = "00000000000000000000",
        AutoDiscovery = DevOpsAutoDiscovery.Disabled,
        TopLevelInventoryList =
        {
        "org1","org2"
        },
    },
};
ArmOperation<DevOpsConfigurationResource> lro = await devOpsConfiguration.CreateOrUpdateAsync(WaitUntil.Completed, data);
DevOpsConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevOpsConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
