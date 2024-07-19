using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConfidentialLedger.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ConfidentialLedger;

// Generated from example definition: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/examples/ManagedCCF_Update.json
// this example is just showing the usage of "ManagedCCF_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedCcfResource created on azure
// for more information of creating ManagedCcfResource, please refer to the document of ManagedCcfResource
string subscriptionId = "0000000-0000-0000-0000-000000000001";
string resourceGroupName = "DummyResourceGroupName";
string appName = "DummyMccfAppName";
ResourceIdentifier managedCcfResourceId = ManagedCcfResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, appName);
ManagedCcfResource managedCcf = client.GetManagedCcfResource(managedCcfResourceId);

// invoke the operation
ManagedCcfData data = new ManagedCcfData(new AzureLocation("EastUS"))
{
    Properties = new ManagedCcfProperties()
    {
        DeploymentType = new ConfidentialLedgerDeploymentType()
        {
            LanguageRuntime = ConfidentialLedgerLanguageRuntime.CPP,
            AppSourceUri = new Uri("https://myaccount.blob.core.windows.net/storage/mccfsource?sv=2022-02-11%st=2022-03-11"),
        },
    },
    Tags =
    {
    ["additionalProps1"] = "additional properties",
    },
};
ArmOperation<ManagedCcfResource> lro = await managedCcf.UpdateAsync(WaitUntil.Completed, data);
ManagedCcfResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedCcfData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
