using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/Shares_Create.json
// this example is just showing the usage of "Shares_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataShareResource created on azure
// for more information of creating DataShareResource, please refer to the document of DataShareResource
string subscriptionId = "12345678-1234-1234-12345678abc";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareName = "Share1";
ResourceIdentifier dataShareResourceId = DataShareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName);
DataShareResource dataShare = client.GetDataShareResource(dataShareResourceId);

// invoke the operation
DataShareData data = new DataShareData()
{
    Description = "share description",
    ShareKind = DataShareKind.CopyBased,
    Terms = "Confidential",
};
ArmOperation<DataShareResource> lro = await dataShare.UpdateAsync(WaitUntil.Completed, data);
DataShareResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataShareData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
