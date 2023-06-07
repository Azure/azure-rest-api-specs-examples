using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataBox;
using Azure.ResourceManager.DataBox.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsCreate.json
// this example is just showing the usage of "Jobs_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "YourSubscriptionId";
string resourceGroupName = "YourResourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataBoxJobResource
DataBoxJobCollection collection = resourceGroupResource.GetDataBoxJobs();

// invoke the operation
string jobName = "TestJobName1";
DataBoxJobData data = new DataBoxJobData(new AzureLocation("westus"), DataBoxJobTransferType.ImportToAzure, new DataBoxSku(DataBoxSkuName.DataBox))
{
    Details = new DataBoxJobDetails(new DataBoxContactDetails("XXXX XXXX", "0000000000", new string[]
{
"xxxx@xxxx.xxx"
})
    {
        PhoneExtension = "",
    })
    {
        ShippingAddress = new DataBoxShippingAddress("XXXX XXXX", "XX", "00000")
        {
            StreetAddress2 = "XXXX XXXX",
            City = "XXXX XXXX",
            StateOrProvince = "XX",
            CompanyName = "XXXX XXXX",
            AddressType = DataBoxShippingAddressType.Commercial,
        },
        DataImportDetails =
        {
        new DataImportDetails(new DataBoxStorageAccountDetails(new ResourceIdentifier("/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName")))
        },
    },
};
ArmOperation<DataBoxJobResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, jobName, data);
DataBoxJobResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
