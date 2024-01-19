using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PlaywrightTesting;
using Azure.ResourceManager.PlaywrightTesting.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_Delete.json
// this example is just showing the usage of "Accounts_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PlaywrightTestingAccountResource created on azure
// for more information of creating PlaywrightTestingAccountResource, please refer to the document of PlaywrightTestingAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "dummyrg";
string name = "myPlaywrightAccount";
ResourceIdentifier playwrightTestingAccountResourceId = PlaywrightTestingAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
PlaywrightTestingAccountResource playwrightTestingAccount = client.GetPlaywrightTestingAccountResource(playwrightTestingAccountResourceId);

// invoke the operation
await playwrightTestingAccount.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
