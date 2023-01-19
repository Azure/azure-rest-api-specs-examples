using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SqlVirtualMachine;
using Azure.ResourceManager.SqlVirtualMachine.Models;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/GetSqlVirtualMachine.json
// this example is just showing the usage of "SqlVirtualMachines_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SqlVmResource
SqlVmCollection collection = resourceGroupResource.GetSqlVms();

// invoke the operation
string sqlVmName = "testvm";
bool result = await collection.ExistsAsync(sqlVmName);

Console.WriteLine($"Succeeded: {result}");
