using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.GraphServices;
using Azure.ResourceManager.GraphServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_Delete.json
// this example is just showing the usage of "Accounts_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GraphServicesAccountResource created on azure
// for more information of creating GraphServicesAccountResource, please refer to the document of GraphServicesAccountResource
string subscriptionId = "11111111-aaaa-1111-bbbb-111111111111";
string resourceGroupName = "testResourceGroupGRAM";
string resourceName = "11111111-aaaa-1111-bbbb-111111111111";
ResourceIdentifier graphServicesAccountResourceId = GraphServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
GraphServicesAccountResource graphServicesAccountResource = client.GetGraphServicesAccountResource(graphServicesAccountResourceId);

// invoke the operation
await graphServicesAccountResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
