using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/GetAlertSubscriptionLocation_example.json
// this example is just showing the usage of "Alerts_GetSubscriptionLevel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityCenterLocationResource created on azure
// for more information of creating SecurityCenterLocationResource, please refer to the document of SecurityCenterLocationResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
AzureLocation ascLocation = new AzureLocation("westeurope");
ResourceIdentifier securityCenterLocationResourceId = SecurityCenterLocationResource.CreateResourceIdentifier(subscriptionId, ascLocation);
SecurityCenterLocationResource securityCenterLocation = client.GetSecurityCenterLocationResource(securityCenterLocationResourceId);

// get the collection of this SubscriptionSecurityAlertResource
SubscriptionSecurityAlertCollection collection = securityCenterLocation.GetSubscriptionSecurityAlerts();

// invoke the operation
string alertName = "2518770965529163669_F144EE95-A3E5-42DA-A279-967D115809AA";
bool result = await collection.ExistsAsync(alertName);

Console.WriteLine($"Succeeded: {result}");
