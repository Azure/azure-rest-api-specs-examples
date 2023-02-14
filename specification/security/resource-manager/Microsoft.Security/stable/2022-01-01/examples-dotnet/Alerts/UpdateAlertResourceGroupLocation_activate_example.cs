using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/UpdateAlertResourceGroupLocation_activate_example.json
// this example is just showing the usage of "Alerts_UpdateResourceGroupLevelStateToActivate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupSecurityAlertResource created on azure
// for more information of creating ResourceGroupSecurityAlertResource, please refer to the document of ResourceGroupSecurityAlertResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "myRg2";
AzureLocation ascLocation = new AzureLocation("westeurope");
string alertName = "2518765996949954086_2325cf9e-42a2-4f72-ae7f-9b863cba2d22";
ResourceIdentifier resourceGroupSecurityAlertResourceId = ResourceGroupSecurityAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ascLocation, alertName);
ResourceGroupSecurityAlertResource resourceGroupSecurityAlert = client.GetResourceGroupSecurityAlertResource(resourceGroupSecurityAlertResourceId);

// invoke the operation
await resourceGroupSecurityAlert.ActivateAsync();

Console.WriteLine($"Succeeded");
