using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/DeleteTagsSubscription.json
// this example is just showing the usage of "Tags_DeleteAtScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TagResource created on azure
// for more information of creating TagResource, please refer to the document of TagResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000";
ResourceIdentifier tagResourceId = TagResource.CreateResourceIdentifier(scope);
TagResource tagResource = client.GetTagResource(tagResourceId);

// invoke the operation
await tagResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
