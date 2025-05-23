using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsDelete.json
// this example is just showing the usage of "TemplateSpecVersions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TemplateSpecVersionResource created on azure
// for more information of creating TemplateSpecVersionResource, please refer to the document of TemplateSpecVersionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "templateSpecRG";
string templateSpecName = "simpleTemplateSpec";
string templateSpecVersion = "v1.0";
ResourceIdentifier templateSpecVersionResourceId = TemplateSpecVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, templateSpecName, templateSpecVersion);
TemplateSpecVersionResource templateSpecVersion0 = client.GetTemplateSpecVersionResource(templateSpecVersionResourceId);

// invoke the operation
await templateSpecVersion0.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
