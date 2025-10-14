using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageMover;

// Generated from example definition: 2025-07-01/JobRuns_List.json
// this example is just showing the usage of "JobRun_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this JobDefinitionResource created on azure
// for more information of creating JobDefinitionResource, please refer to the document of JobDefinitionResource
string subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
string resourceGroupName = "examples-rg";
string storageMoverName = "examples-storageMoverName";
string projectName = "examples-projectName";
string jobDefinitionName = "examples-jobDefinitionName";
ResourceIdentifier jobDefinitionResourceId = JobDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageMoverName, projectName, jobDefinitionName);
JobDefinitionResource jobDefinition = client.GetJobDefinitionResource(jobDefinitionResourceId);

// get the collection of this JobRunResource
JobRunCollection collection = jobDefinition.GetJobRuns();

// invoke the operation and iterate over the result
await foreach (JobRunResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    JobRunData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
