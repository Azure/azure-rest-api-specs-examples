using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataMigration;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/Services_CreateOrUpdate.json
// this example is just showing the usage of "Services_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
string groupName = "DmsSdkRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, groupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataMigrationServiceResource
DataMigrationServiceCollection collection = resourceGroupResource.GetDataMigrationServices();

// invoke the operation
string serviceName = "DmsSdkService";
DataMigrationServiceData data = new DataMigrationServiceData(new AzureLocation("southcentralus"))
{
    Sku = new ServiceSku()
    {
        Name = "Basic_1vCore",
    },
    VirtualSubnetId = "/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkTestNetwork/providers/Microsoft.Network/virtualNetworks/DmsSdkTestNetwork/subnets/default",
};
ArmOperation<DataMigrationServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serviceName, data);
DataMigrationServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataMigrationServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
