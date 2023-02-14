using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/GetAdaptiveApplicationControlsGroup_example.json
// this example is just showing the usage of "AdaptiveApplicationControls_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityCenterLocationResource created on azure
// for more information of creating SecurityCenterLocationResource, please refer to the document of SecurityCenterLocationResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
AzureLocation ascLocation = new AzureLocation("centralus");
ResourceIdentifier securityCenterLocationResourceId = SecurityCenterLocationResource.CreateResourceIdentifier(subscriptionId, ascLocation);
SecurityCenterLocationResource securityCenterLocation = client.GetSecurityCenterLocationResource(securityCenterLocationResourceId);

// get the collection of this AdaptiveApplicationControlGroupResource
AdaptiveApplicationControlGroupCollection collection = securityCenterLocation.GetAdaptiveApplicationControlGroups();

// invoke the operation
string groupName = "ERELGROUP1";
bool result = await collection.ExistsAsync(groupName);

Console.WriteLine($"Succeeded: {result}");
