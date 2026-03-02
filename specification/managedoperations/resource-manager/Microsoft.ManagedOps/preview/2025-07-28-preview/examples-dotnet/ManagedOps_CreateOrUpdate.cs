using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedOps.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagedOps;

// Generated from example definition: 2025-07-28-preview/ManagedOps_CreateOrUpdate.json
// this example is just showing the usage of "ManagedOp_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "11809CA1-E126-4017-945E-AA795CD5C5A9";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ManagedOpResource
ManagedOpCollection collection = subscriptionResource.GetManagedOps();

// invoke the operation
string managedOpsName = "default";
ManagedOpData data = new ManagedOpData
{
    Properties = new ManagedOpsProperties(new ManagedOpsDesiredConfiguration(new ResourceIdentifier("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/00000000-0000-0000-0000-000000000000-Default"), new ResourceIdentifier("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/example"), new ResourceIdentifier("/subscriptions/11809CA1-E126-4017-945E-AA795CD5C5A9/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myManagedIdentity"))),
};
ArmOperation<ManagedOpResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, managedOpsName, data);
ManagedOpResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedOpData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
