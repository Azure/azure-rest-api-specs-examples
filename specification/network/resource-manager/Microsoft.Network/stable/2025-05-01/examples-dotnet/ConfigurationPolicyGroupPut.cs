using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ConfigurationPolicyGroupPut.json
// this example is just showing the usage of "ConfigurationPolicyGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnServerConfigurationResource created on azure
// for more information of creating VpnServerConfigurationResource, please refer to the document of VpnServerConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string vpnServerConfigurationName = "vpnServerConfiguration1";
ResourceIdentifier vpnServerConfigurationResourceId = VpnServerConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vpnServerConfigurationName);
VpnServerConfigurationResource vpnServerConfiguration = client.GetVpnServerConfigurationResource(vpnServerConfigurationResourceId);

// get the collection of this VpnServerConfigurationPolicyGroupResource
VpnServerConfigurationPolicyGroupCollection collection = vpnServerConfiguration.GetVpnServerConfigurationPolicyGroups();

// invoke the operation
string configurationPolicyGroupName = "policyGroup1";
VpnServerConfigurationPolicyGroupData data = new VpnServerConfigurationPolicyGroupData
{
    IsDefault = true,
    Priority = 0,
    PolicyMembers = {new VpnServerConfigurationPolicyGroupMember
    {
    Name = "policy1",
    AttributeType = VpnPolicyMemberAttributeType.RadiusAzureGroupId,
    AttributeValue = "6ad1bd08",
    }, new VpnServerConfigurationPolicyGroupMember
    {
    Name = "policy2",
    AttributeType = VpnPolicyMemberAttributeType.CertificateGroupId,
    AttributeValue = "red.com",
    }},
};
ArmOperation<VpnServerConfigurationPolicyGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationPolicyGroupName, data);
VpnServerConfigurationPolicyGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VpnServerConfigurationPolicyGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
