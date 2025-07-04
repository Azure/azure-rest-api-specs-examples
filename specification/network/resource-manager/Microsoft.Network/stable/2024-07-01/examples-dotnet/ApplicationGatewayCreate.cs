using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ApplicationGatewayCreate.json
// this example is just showing the usage of "ApplicationGateways_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ApplicationGatewayResource
ApplicationGatewayCollection collection = resourceGroupResource.GetApplicationGateways();

// invoke the operation
string applicationGatewayName = "appgw";
ApplicationGatewayData data = new ApplicationGatewayData
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1")] = new UserAssignedIdentity()
        },
    },
    Sku = new ApplicationGatewaySku
    {
        Name = ApplicationGatewaySkuName.StandardV2,
        Tier = ApplicationGatewayTier.StandardV2,
        Capacity = 3,
    },
    GatewayIPConfigurations = {new ApplicationGatewayIPConfiguration
    {
    SubnetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/appgwsubnet"),
    Name = "appgwipc",
    }},
    TrustedRootCertificates = {new ApplicationGatewayTrustedRootCertificate
    {
    Data = BinaryData.FromObjectAsJson("****"),
    Name = "rootcert",
    }, new ApplicationGatewayTrustedRootCertificate
    {
    KeyVaultSecretId = "https://kv/secret",
    Name = "rootcert1",
    }},
    TrustedClientCertificates = {new ApplicationGatewayTrustedClientCertificate
    {
    Data = BinaryData.FromObjectAsJson("****"),
    Name = "clientcert",
    }},
    SslCertificates = {new ApplicationGatewaySslCertificate
    {
    Data = BinaryData.FromObjectAsJson("****"),
    Password = "****",
    Name = "sslcert",
    }, new ApplicationGatewaySslCertificate
    {
    KeyVaultSecretId = "https://kv/secret",
    Name = "sslcert2",
    }},
    FrontendIPConfigurations = {new ApplicationGatewayFrontendIPConfiguration
    {
    PublicIPAddressId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/appgwpip"),
    Name = "appgwfip",
    }},
    FrontendPorts = {new ApplicationGatewayFrontendPort
    {
    Port = 443,
    Name = "appgwfp",
    }, new ApplicationGatewayFrontendPort
    {
    Port = 80,
    Name = "appgwfp80",
    }},
    BackendAddressPools = {new ApplicationGatewayBackendAddressPool
    {
    BackendAddresses = {new ApplicationGatewayBackendAddress
    {
    IPAddress = "10.0.1.1",
    }, new ApplicationGatewayBackendAddress
    {
    IPAddress = "10.0.1.2",
    }},
    Name = "appgwpool",
    }, new ApplicationGatewayBackendAddressPool
    {
    BackendAddresses = {new ApplicationGatewayBackendAddress
    {
    IPAddress = "10.0.0.1",
    }, new ApplicationGatewayBackendAddress
    {
    IPAddress = "10.0.0.2",
    }},
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool1"),
    Name = "appgwpool1",
    }},
    BackendHttpSettingsCollection = {new ApplicationGatewayBackendHttpSettings
    {
    Port = 80,
    Protocol = ApplicationGatewayProtocol.Http,
    CookieBasedAffinity = ApplicationGatewayCookieBasedAffinity.Disabled,
    RequestTimeoutInSeconds = 30,
    Name = "appgwbhs",
    }},
    HttpListeners = {new ApplicationGatewayHttpListener
    {
    FrontendIPConfigurationId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
    FrontendPortId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp"),
    Protocol = ApplicationGatewayProtocol.Https,
    SslCertificateId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert"),
    SslProfileId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1"),
    RequireServerNameIndication = false,
    Name = "appgwhl",
    }, new ApplicationGatewayHttpListener
    {
    FrontendIPConfigurationId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip"),
    FrontendPortId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80"),
    Protocol = ApplicationGatewayProtocol.Http,
    Name = "appgwhttplistener",
    }},
    SslProfiles = {new ApplicationGatewaySslProfile
    {
    TrustedClientCertificates = {new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert"),
    }},
    SslPolicy = new ApplicationGatewaySslPolicy
    {
    PolicyType = ApplicationGatewaySslPolicyType.Custom,
    CipherSuites = {ApplicationGatewaySslCipherSuite.TlsECDiffieHellmanRsaWithAes128CbcSha256},
    MinProtocolVersion = ApplicationGatewaySslProtocol.Tls1_1,
    },
    ClientAuthConfiguration = new ApplicationGatewayClientAuthConfiguration
    {
    VerifyClientCertIssuerDN = true,
    },
    Name = "sslProfile1",
    }},
    RequestRoutingRules = {new ApplicationGatewayRequestRoutingRule
    {
    RuleType = ApplicationGatewayRequestRoutingRuleType.Basic,
    Priority = 10,
    BackendAddressPoolId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool"),
    BackendHttpSettingsId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs"),
    HttpListenerId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl"),
    RewriteRuleSetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1"),
    Name = "appgwrule",
    }},
    RewriteRuleSets = {new ApplicationGatewayRewriteRuleSet
    {
    RewriteRules = {new ApplicationGatewayRewriteRule
    {
    Name = "Set X-Forwarded-For",
    RuleSequence = 102,
    Conditions = {new ApplicationGatewayRewriteRuleCondition
    {
    Variable = "http_req_Authorization",
    Pattern = "^Bearer",
    IgnoreCase = true,
    Negate = false,
    }},
    ActionSet = new ApplicationGatewayRewriteRuleActionSet
    {
    RequestHeaderConfigurations = {new ApplicationGatewayHeaderConfiguration
    {
    HeaderName = "X-Forwarded-For",
    HeaderValue = "{var_add_x_forwarded_for_proxy}",
    }},
    ResponseHeaderConfigurations = {new ApplicationGatewayHeaderConfiguration
    {
    HeaderName = "Strict-Transport-Security",
    HeaderValue = "max-age=31536000",
    }},
    UrlConfiguration = new ApplicationGatewayUrlConfiguration
    {
    ModifiedPath = "/abc",
    },
    },
    }},
    Name = "rewriteRuleSet1",
    }},
    GlobalConfiguration = new ApplicationGatewayGlobalConfiguration
    {
        EnableRequestBuffering = true,
        EnableResponseBuffering = true,
    },
    Location = new AzureLocation("eastus"),
};
ArmOperation<ApplicationGatewayResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, applicationGatewayName, data);
ApplicationGatewayResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApplicationGatewayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
