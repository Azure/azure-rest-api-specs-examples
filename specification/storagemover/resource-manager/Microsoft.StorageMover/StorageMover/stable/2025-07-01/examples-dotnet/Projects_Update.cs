using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageMover.Models;
using Azure.ResourceManager.StorageMover;

// Generated from example definition: 2025-07-01/Projects_Update.json
// this example is just showing the usage of "Project_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageMoverProjectResource created on azure
// for more information of creating StorageMoverProjectResource, please refer to the document of StorageMoverProjectResource
string subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
string resourceGroupName = "examples-rg";
string storageMoverName = "examples-storageMoverName";
string projectName = "examples-projectName";
ResourceIdentifier storageMoverProjectResourceId = StorageMoverProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageMoverName, projectName);
StorageMoverProjectResource storageMoverProject = client.GetStorageMoverProjectResource(storageMoverProjectResourceId);

// invoke the operation
StorageMoverProjectPatch patch = new StorageMoverProjectPatch();
StorageMoverProjectResource result = await storageMoverProject.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageMoverProjectData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
