using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/DeleteAdaptiveApplicationControls_example.json
// this example is just showing the usage of "AdaptiveApplicationControls_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AdaptiveApplicationControlGroupResource created on azure
// for more information of creating AdaptiveApplicationControlGroupResource, please refer to the document of AdaptiveApplicationControlGroupResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
AzureLocation ascLocation = new AzureLocation("centralus");
string groupName = "GROUP1";
ResourceIdentifier adaptiveApplicationControlGroupResourceId = AdaptiveApplicationControlGroupResource.CreateResourceIdentifier(subscriptionId, ascLocation, groupName);
AdaptiveApplicationControlGroupResource adaptiveApplicationControlGroup = client.GetAdaptiveApplicationControlGroupResource(adaptiveApplicationControlGroupResourceId);

// invoke the operation
await adaptiveApplicationControlGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
