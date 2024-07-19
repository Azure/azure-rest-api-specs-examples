using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedInstanceCreateMin.json
// this example is just showing the usage of "ManagedInstances_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ManagedInstanceResource
ManagedInstanceCollection collection = resourceGroupResource.GetManagedInstances();

// invoke the operation
string managedInstanceName = "testinstance";
ManagedInstanceData data = new ManagedInstanceData(new AzureLocation("Japan East"))
{
    Sku = new SqlSku("GP_Gen4")
    {
        Tier = "GeneralPurpose",
    },
    AdministratorLogin = "dummylogin",
    AdministratorLoginPassword = "PLACEHOLDER",
    SubnetId = new ResourceIdentifier("/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
    LicenseType = ManagedInstanceLicenseType.LicenseIncluded,
    VCores = 8,
    StorageSizeInGB = 1024,
};
ArmOperation<ManagedInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, managedInstanceName, data);
ManagedInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
