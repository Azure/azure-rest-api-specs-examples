using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Install_SingleServer.json
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
SapVirtualInstanceData data = new SapVirtualInstanceData(new AzureLocation("eastus2"), SapEnvironmentType.NonProd, SapProductType.S4Hana, new DeploymentWithOSConfiguration()
{
    AppLocation = new AzureLocation("eastus"),
    InfrastructureConfiguration = new SingleServerConfiguration("test-rg", new ResourceIdentifier("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/testsubnet"), new SapVirtualMachineConfiguration("Standard_E32ds_v4", new SapImageReference()
    {
        Publisher = "SUSE",
        Offer = "SLES-SAP",
        Sku = "12-sp4-gen2",
        Version = "2022.02.01",
    }, new SapOSProfile()
    {
        AdminUsername = "azureappadmin",
        OSConfiguration = new SapLinuxConfiguration()
        {
            DisablePasswordAuthentication = true,
            SshKeyPair = new SapSshKeyPair()
            {
                PublicKey = "{{sshkey}}",
                PrivateKey = "{{privateKey}}",
            },
        },
    })),
    SoftwareConfiguration = new SapInstallWithoutOSConfigSoftwareConfiguration(new Uri("https://teststorageaccount.blob.core.windows.net/sapbits/sapfiles/boms/S41909SPS03_v0011ms/S41909SPS03_v0011ms.yaml"), "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorageaccount", "SAP S/4HANA 1909 SPS 03"),
    OSSapConfiguration = new OSSapConfiguration()
    {
        SapFqdn = "sap.bpaas.com",
    },
})
{
    Tags =
    {
    },
};
ArmOperation<SapVirtualInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sapVirtualInstanceName, data);
SapVirtualInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapVirtualInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
