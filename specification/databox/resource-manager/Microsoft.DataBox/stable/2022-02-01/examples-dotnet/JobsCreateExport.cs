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

// Generated from example definition: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsCreateExport.json
// this example is just showing the usage of "Jobs_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
string resourceGroupName = "SdkRg8091";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataBoxJobResource
DataBoxJobCollection collection = resourceGroupResource.GetDataBoxJobs();

// invoke the operation
string jobName = "SdkJob6429";
DataBoxJobData data = new DataBoxJobData(new AzureLocation("westus"), DataBoxJobTransferType.ExportFromAzure, new DataBoxSku(DataBoxSkuName.DataBox))
{
    Details = new DataBoxJobDetails(new DataBoxContactDetails("Public SDK Test", "1234567890", new string[]
{
"testing@microsoft.com"
})
    {
        PhoneExtension = "1234",
    })
    {
        ShippingAddress = new DataBoxShippingAddress("16 TOWNSEND ST", "US", "94107")
        {
            StreetAddress2 = "Unit 1",
            City = "San Francisco",
            StateOrProvince = "CA",
            CompanyName = "Microsoft",
            AddressType = DataBoxShippingAddressType.Commercial,
        },
        DataExportDetails =
        {
        new DataExportDetails(new TransferConfiguration(TransferConfigurationType.TransferAll)
        {
        TransferAllDetailsInclude = new TransferAllDetails(DataAccountType.StorageAccount)
        {
        TransferAllBlobs = true,
        TransferAllFiles = true,
        },
        },new DataBoxStorageAccountDetails(new ResourceIdentifier("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.Storage/storageAccounts/aaaaaa2")))
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
