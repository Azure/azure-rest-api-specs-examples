using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Support.Models;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateFileForSubscription.json
// this example is just showing the usage of "Files_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportTicketFileResource created on azure
// for more information of creating SupportTicketFileResource, please refer to the document of SupportTicketFileResource
string subscriptionId = "132d901f-189d-4381-9214-fe68e27e05a1";
string fileWorkspaceName = "testworkspace";
string fileName = "test.txt";
ResourceIdentifier supportTicketFileResourceId = SupportTicketFileResource.CreateResourceIdentifier(subscriptionId, fileWorkspaceName, fileName);
SupportTicketFileResource supportTicketFile = client.GetSupportTicketFileResource(supportTicketFileResourceId);

// invoke the operation
SupportFileDetailData data = new SupportFileDetailData
{
    ChunkSize = 41423,
    FileSize = 41423,
    NumberOfChunks = 1,
};
ArmOperation<SupportTicketFileResource> lro = await supportTicketFile.UpdateAsync(WaitUntil.Completed, data);
SupportTicketFileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportFileDetailData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
