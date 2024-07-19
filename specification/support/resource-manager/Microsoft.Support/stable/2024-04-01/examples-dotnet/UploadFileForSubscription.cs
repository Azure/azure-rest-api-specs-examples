using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Support.Models;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/UploadFileForSubscription.json
// this example is just showing the usage of "Files_Upload" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportTicketFileResource created on azure
// for more information of creating SupportTicketFileResource, please refer to the document of SupportTicketFileResource
string subscriptionId = "132d901f-189d-4381-9214-fe68e27e05a1";
string fileWorkspaceName = "testworkspaceName";
string fileName = "test.txt";
ResourceIdentifier supportTicketFileResourceId = SupportTicketFileResource.CreateResourceIdentifier(subscriptionId, fileWorkspaceName, fileName);
SupportTicketFileResource supportTicketFile = client.GetSupportTicketFileResource(supportTicketFileResourceId);

// invoke the operation
UploadFileContent content = new UploadFileContent()
{
    Content = "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAMAAAAoLQ9TAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABd",
    ChunkIndex = 0,
};
await supportTicketFile.UploadAsync(content);

Console.WriteLine($"Succeeded");
