using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GraphServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.GraphServices;

// Generated from example definition: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_Create.json
// this example is just showing the usage of "Accounts_CreateAndUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testResourceGroupGRAM";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this GraphServicesAccountResource
GraphServicesAccountResourceCollection collection = resourceGroupResource.GetGraphServicesAccountResources();

// invoke the operation
string resourceName = "11111111-aaaa-1111-bbbb-1111111111111";
GraphServicesAccountResourceData data = new GraphServicesAccountResourceData(default, new GraphServicesAccountResourceProperties("11111111-aaaa-1111-bbbb-111111111111"));
ArmOperation<GraphServicesAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
GraphServicesAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GraphServicesAccountResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
