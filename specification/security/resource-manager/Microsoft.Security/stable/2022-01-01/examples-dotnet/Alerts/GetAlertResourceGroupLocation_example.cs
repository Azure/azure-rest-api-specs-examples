using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/GetAlertResourceGroupLocation_example.json
// this example is just showing the usage of "Alerts_GetResourceGroupLevel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupSecurityAlertResource created on azure
// for more information of creating ResourceGroupSecurityAlertResource, please refer to the document of ResourceGroupSecurityAlertResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "myRg1";
AzureLocation ascLocation = new AzureLocation("westeurope");
string alertName = "2518298467986649999_4d25bfef-2d77-4a08-adc0-3e35715cc92a";
ResourceIdentifier resourceGroupSecurityAlertResourceId = ResourceGroupSecurityAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ascLocation, alertName);
ResourceGroupSecurityAlertResource resourceGroupSecurityAlert = client.GetResourceGroupSecurityAlertResource(resourceGroupSecurityAlertResourceId);

// invoke the operation
ResourceGroupSecurityAlertResource result = await resourceGroupSecurityAlert.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
