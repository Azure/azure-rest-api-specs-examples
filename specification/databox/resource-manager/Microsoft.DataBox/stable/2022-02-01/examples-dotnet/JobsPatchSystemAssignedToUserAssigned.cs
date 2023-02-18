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

// Generated from example definition: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsPatchSystemAssignedToUserAssigned.json
// this example is just showing the usage of "Jobs_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxJobResource created on azure
// for more information of creating DataBoxJobResource, please refer to the document of DataBoxJobResource
string subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
string resourceGroupName = "SdkRg9765";
string jobName = "SdkJob2965";
ResourceIdentifier dataBoxJobResourceId = DataBoxJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
DataBoxJobResource dataBoxJob = client.GetDataBoxJobResource(dataBoxJobResourceId);

// invoke the operation
DataBoxJobPatch patch = new DataBoxJobPatch()
{
    Identity = new ManagedServiceIdentity("SystemAssigned,UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity")] = new UserAssignedIdentity(),
        },
    },
    Details = new UpdateJobDetails()
    {
        KeyEncryptionKey = new DataBoxKeyEncryptionKey(DataBoxKeyEncryptionKeyType.CustomerManaged)
        {
            ManagedIdentity = new DataBoxManagedIdentity()
            {
                IdentityType = "UserAssigned",
                UserAssignedResourceId = new ResourceIdentifier("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity"),
            },
            KekUri = new Uri("https://sdkkeyvault.vault.azure.net/keys/SSDKEY/"),
            KekVaultResourceId = new ResourceIdentifier("/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.KeyVault/vaults/SDKKeyVault"),
        },
    },
};
ArmOperation<DataBoxJobResource> lro = await dataBoxJob.UpdateAsync(WaitUntil.Completed, patch);
DataBoxJobResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
