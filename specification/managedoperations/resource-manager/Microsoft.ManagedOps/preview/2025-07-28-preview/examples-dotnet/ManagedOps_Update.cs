using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedOps.Models;
using Azure.ResourceManager.ManagedOps;

// Generated from example definition: 2025-07-28-preview/ManagedOps_Update.json
// this example is just showing the usage of "ManagedOp_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedOpResource created on azure
// for more information of creating ManagedOpResource, please refer to the document of ManagedOpResource
string subscriptionId = "11809CA1-E126-4017-945E-AA795CD5C5A9";
string managedOpsName = "default";
ResourceIdentifier managedOpResourceId = ManagedOpResource.CreateResourceIdentifier(subscriptionId, managedOpsName);
ManagedOpResource managedOp = client.GetManagedOpResource(managedOpResourceId);

// invoke the operation
ManagedOpPatch patch = new ManagedOpPatch();
ArmOperation<ManagedOpResource> lro = await managedOp.UpdateAsync(WaitUntil.Completed, patch);
ManagedOpResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedOpData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
