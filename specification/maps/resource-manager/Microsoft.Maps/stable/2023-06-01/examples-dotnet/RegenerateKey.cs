using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Maps.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Maps;

// Generated from example definition: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/RegenerateKey.json
// this example is just showing the usage of "Accounts_RegenerateKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MapsAccountResource created on azure
// for more information of creating MapsAccountResource, please refer to the document of MapsAccountResource
string subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
string resourceGroupName = "myResourceGroup";
string accountName = "myMapsAccount";
ResourceIdentifier mapsAccountResourceId = MapsAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
MapsAccountResource mapsAccount = client.GetMapsAccountResource(mapsAccountResourceId);

// invoke the operation
MapsKeySpecification keySpecification = new MapsKeySpecification(MapsKeyType.Primary);
MapsAccountKeys result = await mapsAccount.RegenerateKeysAsync(keySpecification);

Console.WriteLine($"Succeeded: {result}");
