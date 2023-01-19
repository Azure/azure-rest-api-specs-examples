using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ExtendedLocations;
using Azure.ResourceManager.ExtendedLocations.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsDelete.json
// this example is just showing the usage of "CustomLocations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CustomLocationResource created on azure
// for more information of creating CustomLocationResource, please refer to the document of CustomLocationResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
string resourceGroupName = "testresourcegroup";
string resourceName = "customLocation01";
ResourceIdentifier customLocationResourceId = CustomLocationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
CustomLocationResource customLocation = client.GetCustomLocationResource(customLocationResourceId);

// invoke the operation
await customLocation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
