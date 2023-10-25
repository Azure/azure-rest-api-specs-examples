using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;
using Azure.ResourceManager.Support.Models;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/UploadFile.json
// this example is just showing the usage of "FilesNoSubscription_Upload" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SupportTicketNoSubFileResource created on azure
// for more information of creating SupportTicketNoSubFileResource, please refer to the document of SupportTicketNoSubFileResource
string fileWorkspaceName = "testworkspaceName";
string fileName = "test.txt";
ResourceIdentifier supportTicketNoSubFileResourceId = SupportTicketNoSubFileResource.CreateResourceIdentifier(fileWorkspaceName, fileName);
SupportTicketNoSubFileResource supportTicketNoSubFile = client.GetSupportTicketNoSubFileResource(supportTicketNoSubFileResourceId);

// invoke the operation
UploadFileContent content = new UploadFileContent()
{
    Content = "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAMAAAAoLQ9TAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABd",
    ChunkIndex = 0,
};
await supportTicketNoSubFile.UploadAsync(content);

Console.WriteLine($"Succeeded");
