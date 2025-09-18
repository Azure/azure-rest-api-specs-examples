using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Jobs_ListByTarget_MaximumSet_Gen.json
// this example is just showing the usage of "Job_ListByTarget" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this EdgeJobResource
string resourceUri = "gt";
EdgeJobCollection collection = client.GetEdgeJobs(new ResourceIdentifier(resourceUri));

// invoke the operation and iterate over the result
await foreach (EdgeJobResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EdgeJobData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
