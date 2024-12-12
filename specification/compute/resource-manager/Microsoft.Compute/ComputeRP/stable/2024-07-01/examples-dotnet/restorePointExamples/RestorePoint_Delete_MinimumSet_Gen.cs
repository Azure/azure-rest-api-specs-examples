using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/restorePointExamples/RestorePoint_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "RestorePoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorePointResource created on azure
// for more information of creating RestorePointResource, please refer to the document of RestorePointResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string restorePointGroupName = "aaaaaaaaaaaaaaaaa";
string restorePointName = "aaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier restorePointResourceId = RestorePointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, restorePointGroupName, restorePointName);
RestorePointResource restorePoint = client.GetRestorePointResource(restorePointResourceId);

// invoke the operation
await restorePoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
