using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentCreate_example.json
// this example is just showing the usage of "CustomEntityStoreAssignments_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CustomEntityStoreAssignmentResource created on azure
// for more information of creating CustomEntityStoreAssignmentResource, please refer to the document of CustomEntityStoreAssignmentResource
string subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
string resourceGroupName = "TestResourceGroup";
string customEntityStoreAssignmentName = "33e7cc6e-a139-4723-a0e5-76993aee0771";
ResourceIdentifier customEntityStoreAssignmentResourceId = CustomEntityStoreAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, customEntityStoreAssignmentName);
CustomEntityStoreAssignmentResource customEntityStoreAssignment = client.GetCustomEntityStoreAssignmentResource(customEntityStoreAssignmentResourceId);

// invoke the operation
CustomEntityStoreAssignmentCreateOrUpdateContent content = new CustomEntityStoreAssignmentCreateOrUpdateContent()
{
    Principal = "aaduser=f3923a3e-ad57-4752-b1a9-fbf3c8e5e082;72f988bf-86f1-41af-91ab-2d7cd011db47",
};
ArmOperation<CustomEntityStoreAssignmentResource> lro = await customEntityStoreAssignment.UpdateAsync(WaitUntil.Completed, content);
CustomEntityStoreAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CustomEntityStoreAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
