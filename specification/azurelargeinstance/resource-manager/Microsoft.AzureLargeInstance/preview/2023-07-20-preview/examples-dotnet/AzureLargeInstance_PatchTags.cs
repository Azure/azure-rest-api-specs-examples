using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LargeInstance;
using Azure.ResourceManager.LargeInstance.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeInstance_PatchTags.json
// this example is just showing the usage of "AzureLargeInstance_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LargeInstanceResource created on azure
// for more information of creating LargeInstanceResource, please refer to the document of LargeInstanceResource
string subscriptionId = "f0f4887f-d13c-4943-a8ba-d7da28d2a3fd";
string resourceGroupName = "myResourceGroup";
string azureLargeInstanceName = "myALInstance";
ResourceIdentifier largeInstanceResourceId = LargeInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureLargeInstanceName);
LargeInstanceResource largeInstance = client.GetLargeInstanceResource(largeInstanceResourceId);

// invoke the operation
LargeInstancePatch patch = new LargeInstancePatch();
LargeInstanceResource result = await largeInstance.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LargeInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
