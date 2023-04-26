using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Create_WithOSConfig_SingleServer.json
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
SapVirtualInstanceData data = new SapVirtualInstanceData(new AzureLocation("westcentralus"), SapEnvironmentType.NonProd, SapProductType.S4Hana, new DeploymentWithOSConfiguration()
{
    AppLocation = new AzureLocation("eastus"),
    InfrastructureConfiguration = new SingleServerConfiguration("X00-RG", new ResourceIdentifier("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"), new SapVirtualMachineConfiguration("Standard_E32ds_v4", new SapImageReference()
    {
        Publisher = "RedHat",
        Offer = "RHEL-SAP",
        Sku = "84sapha-gen2",
        Version = "latest",
    }, new SapOSProfile()
    {
        AdminUsername = "{your-username}",
        OSConfiguration = new SapLinuxConfiguration()
        {
            DisablePasswordAuthentication = true,
            SshKeyPair = new SapSshKeyPair()
            {
                PublicKey = "abc",
                PrivateKey = "xyz",
            },
        },
    }))
    {
        IsSecondaryIPEnabled = true,
        DatabaseType = SapDatabaseType.Hana,
    },
    OSSapConfiguration = new OSSapConfiguration()
    {
        SapFqdn = "xyz.test.com",
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
