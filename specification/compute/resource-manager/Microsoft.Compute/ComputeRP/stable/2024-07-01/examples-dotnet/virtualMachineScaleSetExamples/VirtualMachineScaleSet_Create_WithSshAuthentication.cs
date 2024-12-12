using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSshAuthentication.json
// this example is just showing the usage of "VirtualMachineScaleSets_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VirtualMachineScaleSetResource
VirtualMachineScaleSetCollection collection = resourceGroupResource.GetVirtualMachineScaleSets();

// invoke the operation
string virtualMachineScaleSetName = "{vmss-name}";
VirtualMachineScaleSetData data = new VirtualMachineScaleSetData(new AzureLocation("westus"))
{
    Sku = new ComputeSku
    {
        Name = "Standard_D1_v2",
        Tier = "Standard",
        Capacity = 3L,
    },
    Properties = new VirtualMachineScaleSetProperties
    {
        UpgradePolicy = new VirtualMachineScaleSetUpgradePolicy
        {
            Mode = VirtualMachineScaleSetUpgradeMode.Manual,
        },
        VirtualMachineProfile = new VirtualMachineScaleSetVmProfile
        {
            OSProfile = new VirtualMachineScaleSetOSProfile
            {
                ComputerNamePrefix = "{vmss-name}",
                AdminUsername = "{your-username}",
                LinuxConfiguration = new LinuxConfiguration
                {
                    IsPasswordAuthenticationDisabled = true,
                    SshPublicKeys = {new SshPublicKeyConfiguration
                    {
                    Path = "/home/{your-username}/.ssh/authorized_keys",
                    KeyData = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1",
                    }},
                },
            },
            StorageProfile = new VirtualMachineScaleSetStorageProfile
            {
                ImageReference = new ImageReference
                {
                    Publisher = "MicrosoftWindowsServer",
                    Offer = "WindowsServer",
                    Sku = "2016-Datacenter",
                    Version = "latest",
                },
                OSDisk = new VirtualMachineScaleSetOSDisk(DiskCreateOptionType.FromImage)
                {
                    Caching = CachingType.ReadWrite,
                    ManagedDisk = new VirtualMachineScaleSetManagedDisk
                    {
                        StorageAccountType = StorageAccountType.StandardLrs,
                    },
                },
            },
            NetworkProfile = new VirtualMachineScaleSetNetworkProfile
            {
                NetworkInterfaceConfigurations = {new VirtualMachineScaleSetNetworkConfiguration("{vmss-name}")
                {
                Primary = true,
                IPConfigurations = {new VirtualMachineScaleSetIPConfiguration("{vmss-name}")
                {
                SubnetId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
                }},
                EnableIPForwarding = true,
                }},
            },
        },
        Overprovision = true,
    },
};
ArmOperation<VirtualMachineScaleSetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, virtualMachineScaleSetName, data);
VirtualMachineScaleSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineScaleSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
