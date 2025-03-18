using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Workloads.Models;
using Azure.ResourceManager.Workloads;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Create_CustomFullResourceNames_HA_AvSet.json
// this example is just showing the usage of "SAPVirtualInstances_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
string resourceGroupName = "test-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SapVirtualInstanceResource
SapVirtualInstanceCollection collection = resourceGroupResource.GetSapVirtualInstances();

// invoke the operation
string sapVirtualInstanceName = "X00";
SapVirtualInstanceData data = new SapVirtualInstanceData(new AzureLocation("westcentralus"), SapEnvironmentType.Prod, SapProductType.S4Hana, new DeploymentWithOSConfiguration
{
    AppLocation = new AzureLocation("eastus"),
    InfrastructureConfiguration = new ThreeTierConfiguration("X00-RG", new CentralServerConfiguration(new ResourceIdentifier("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"), new SapVirtualMachineConfiguration("Standard_E16ds_v4", new SapImageReference
    {
        Publisher = "RedHat",
        Offer = "RHEL-SAP",
        Sku = "84sapha-gen2",
        Version = "latest",
    }, new SapOSProfile
    {
        AdminUsername = "{your-username}",
        OSConfiguration = new SapLinuxConfiguration
        {
            DisablePasswordAuthentication = true,
            SshKeyPair = new SapSshKeyPair
            {
                PublicKey = "abc",
                PrivateKey = "xyz",
            },
        },
    }), 2L), new ApplicationServerConfiguration(new ResourceIdentifier("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"), new SapVirtualMachineConfiguration("Standard_E32ds_v4", new SapImageReference
    {
        Publisher = "RedHat",
        Offer = "RHEL-SAP",
        Sku = "84sapha-gen2",
        Version = "latest",
    }, new SapOSProfile
    {
        AdminUsername = "{your-username}",
        OSConfiguration = new SapLinuxConfiguration
        {
            DisablePasswordAuthentication = true,
            SshKeyPair = new SapSshKeyPair
            {
                PublicKey = "abc",
                PrivateKey = "xyz",
            },
        },
    }), 6L), new DatabaseConfiguration(new ResourceIdentifier("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"), new SapVirtualMachineConfiguration("Standard_M32ts", new SapImageReference
    {
        Publisher = "RedHat",
        Offer = "RHEL-SAP",
        Sku = "84sapha-gen2",
        Version = "latest",
    }, new SapOSProfile
    {
        AdminUsername = "{your-username}",
        OSConfiguration = new SapLinuxConfiguration
        {
            DisablePasswordAuthentication = true,
            SshKeyPair = new SapSshKeyPair
            {
                PublicKey = "abc",
                PrivateKey = "xyz",
            },
        },
    }), 2L)
    {
        DatabaseType = SapDatabaseType.Hana,
    })
    {
        HighAvailabilityType = SapHighAvailabilityType.AvailabilitySet,
        CustomResourceNames = new ThreeTierFullResourceNames
        {
            CentralServer = new CentralServerFullResourceNames
            {
                VirtualMachines = {new VirtualMachineResourceNames
                {
                VmName = "ascsvm",
                HostName = "ascshostName",
                NetworkInterfaces = {new NetworkInterfaceResourceNames
                {
                NetworkInterfaceName = "ascsnic",
                }},
                OSDiskName = "ascsosdisk",
                }, new VirtualMachineResourceNames
                {
                VmName = "ersvm",
                HostName = "ershostName",
                NetworkInterfaces = {new NetworkInterfaceResourceNames
                {
                NetworkInterfaceName = "ersnic",
                }},
                OSDiskName = "ersosdisk",
                }},
                AvailabilitySetName = "csAvSet",
                LoadBalancer = new LoadBalancerResourceNames
                {
                    LoadBalancerName = "ascslb",
                    FrontendIPConfigurationNames = { "ascsip0", "ersip0" },
                    BackendPoolNames = { "ascsBackendPool" },
                    HealthProbeNames = { "ascsHealthProbe", "ersHealthProbe" },
                },
            },
            ApplicationServer = new ApplicationServerFullResourceNames
            {
                VirtualMachines = {new VirtualMachineResourceNames
                {
                VmName = "appvm0",
                HostName = "apphostName0",
                NetworkInterfaces = {new NetworkInterfaceResourceNames
                {
                NetworkInterfaceName = "appnic0",
                }},
                OSDiskName = "app0osdisk",
                DataDiskNames =
                {
                ["default"] = new string[]{"app0disk0"}
                },
                }, new VirtualMachineResourceNames
                {
                VmName = "appvm1",
                HostName = "apphostName1",
                NetworkInterfaces = {new NetworkInterfaceResourceNames
                {
                NetworkInterfaceName = "appnic1",
                }},
                OSDiskName = "app1osdisk",
                DataDiskNames =
                {
                ["default"] = new string[]{"app1disk0"}
                },
                }},
                AvailabilitySetName = "appAvSet",
            },
            DatabaseServer = new DatabaseServerFullResourceNames
            {
                VirtualMachines = {new VirtualMachineResourceNames
                {
                VmName = "dbvmpr",
                HostName = "dbprhostName",
                NetworkInterfaces = {new NetworkInterfaceResourceNames
                {
                NetworkInterfaceName = "dbprnic",
                }},
                OSDiskName = "dbprosdisk",
                DataDiskNames =
                {
                ["hanaData"] = new string[]{"hanadatapr0", "hanadatapr1"},
                ["hanaLog"] = new string[]{"hanalogpr0", "hanalogpr1", "hanalogpr2"},
                ["hanaShared"] = new string[]{"hanasharedpr0", "hanasharedpr1"},
                ["usrSap"] = new string[]{"usrsappr0"}
                },
                }, new VirtualMachineResourceNames
                {
                VmName = "dbvmsr",
                HostName = "dbsrhostName",
                NetworkInterfaces = {new NetworkInterfaceResourceNames
                {
                NetworkInterfaceName = "dbsrnic",
                }},
                OSDiskName = "dbsrosdisk",
                DataDiskNames =
                {
                ["hanaData"] = new string[]{"hanadatasr0", "hanadatasr1"},
                ["hanaLog"] = new string[]{"hanalogsr0", "hanalogsr1", "hanalogsr2"},
                ["hanaShared"] = new string[]{"hanasharedsr0", "hanasharedsr1"},
                ["usrSap"] = new string[]{"usrsapsr0"}
                },
                }},
                AvailabilitySetName = "dbAvSet",
                LoadBalancer = new LoadBalancerResourceNames
                {
                    LoadBalancerName = "dblb",
                    FrontendIPConfigurationNames = { "dbip" },
                    BackendPoolNames = { "dbBackendPool" },
                    HealthProbeNames = { "dbHealthProbe" },
                },
            },
            SharedStorage = new SharedStorageResourceNames
            {
                SharedStorageAccountName = "storageacc",
                SharedStorageAccountPrivateEndPointName = "peForxNFS",
            },
        },
    },
    OSSapConfiguration = new OSSapConfiguration
    {
        SapFqdn = "xyz.test.com",
    },
})
{
    Tags = { },
};
ArmOperation<SapVirtualInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sapVirtualInstanceName, data);
SapVirtualInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapVirtualInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
