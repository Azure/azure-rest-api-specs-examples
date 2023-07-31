using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/DscpConfigurationCreate.json
// this example is just showing the usage of "DscpConfiguration_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DscpConfigurationResource created on azure
// for more information of creating DscpConfigurationResource, please refer to the document of DscpConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string dscpConfigurationName = "mydscpconfig";
ResourceIdentifier dscpConfigurationResourceId = DscpConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dscpConfigurationName);
DscpConfigurationResource dscpConfiguration = client.GetDscpConfigurationResource(dscpConfigurationResourceId);

// invoke the operation
DscpConfigurationData data = new DscpConfigurationData()
{
    QosDefinitionCollection =
    {
    new DscpQosDefinition()
    {
    Markings =
    {
    1
    },
    SourceIPRanges =
    {
    new QosIPRange()
    {
    StartIP = "127.0.0.1",
    EndIP = "127.0.0.2",
    }
    },
    DestinationIPRanges =
    {
    new QosIPRange()
    {
    StartIP = "127.0.10.1",
    EndIP = "127.0.10.2",
    }
    },
    SourcePortRanges =
    {
    new QosPortRange()
    {
    Start = 10,
    End = 11,
    },new QosPortRange()
    {
    Start = 20,
    End = 21,
    }
    },
    DestinationPortRanges =
    {
    new QosPortRange()
    {
    Start = 15,
    End = 15,
    }
    },
    Protocol = ProtocolType.Tcp,
    },new DscpQosDefinition()
    {
    Markings =
    {
    2
    },
    SourceIPRanges =
    {
    new QosIPRange()
    {
    StartIP = "12.0.0.1",
    EndIP = "12.0.0.2",
    }
    },
    DestinationIPRanges =
    {
    new QosIPRange()
    {
    StartIP = "12.0.10.1",
    EndIP = "12.0.10.2",
    }
    },
    SourcePortRanges =
    {
    new QosPortRange()
    {
    Start = 11,
    End = 12,
    }
    },
    DestinationPortRanges =
    {
    new QosPortRange()
    {
    Start = 51,
    End = 52,
    }
    },
    Protocol = ProtocolType.Udp,
    }
    },
    Location = new AzureLocation("eastus"),
};
ArmOperation<DscpConfigurationResource> lro = await dscpConfiguration.UpdateAsync(WaitUntil.Completed, data);
DscpConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DscpConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
