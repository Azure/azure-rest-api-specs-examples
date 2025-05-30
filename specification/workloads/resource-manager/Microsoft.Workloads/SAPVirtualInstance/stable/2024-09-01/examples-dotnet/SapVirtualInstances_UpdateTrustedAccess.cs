using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadsSapVirtualInstance.Models;
using Azure.ResourceManager.WorkloadsSapVirtualInstance;

// Generated from example definition: 2024-09-01/SapVirtualInstances_UpdateTrustedAccess.json
// this example is just showing the usage of "SAPVirtualInstance_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapVirtualInstanceResource created on azure
// for more information of creating SapVirtualInstanceResource, please refer to the document of SapVirtualInstanceResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
string resourceGroupName = "test-rg";
string sapVirtualInstanceName = "X00";
ResourceIdentifier sapVirtualInstanceResourceId = SapVirtualInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sapVirtualInstanceName);
SapVirtualInstanceResource sapVirtualInstance = client.GetSapVirtualInstanceResource(sapVirtualInstanceResourceId);

// invoke the operation
SapVirtualInstancePatch patch = new SapVirtualInstancePatch
{
    Tags =
    {
    ["key1"] = "svi1"
    },
    Identity = new SapVirtualInstanceIdentity(SapVirtualInstanceIdentityType.None),
    UpdateSapVirtualInstanceManagedResourcesNetworkAccessType = ManagedResourcesNetworkAccessType.Private,
};
ArmOperation<SapVirtualInstanceResource> lro = await sapVirtualInstance.UpdateAsync(WaitUntil.Completed, patch);
SapVirtualInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapVirtualInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
