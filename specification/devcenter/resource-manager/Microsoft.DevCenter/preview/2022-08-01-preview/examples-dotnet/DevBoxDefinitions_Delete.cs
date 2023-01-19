using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/DevBoxDefinitions_Delete.json
// this example is just showing the usage of "DevBoxDefinitions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevBoxDefinitionResource created on azure
// for more information of creating DevBoxDefinitionResource, please refer to the document of DevBoxDefinitionResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
string devBoxDefinitionName = "WebDevBox";
ResourceIdentifier devBoxDefinitionResourceId = DevBoxDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName, devBoxDefinitionName);
DevBoxDefinitionResource devBoxDefinition = client.GetDevBoxDefinitionResource(devBoxDefinitionResourceId);

// invoke the operation
await devBoxDefinition.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
