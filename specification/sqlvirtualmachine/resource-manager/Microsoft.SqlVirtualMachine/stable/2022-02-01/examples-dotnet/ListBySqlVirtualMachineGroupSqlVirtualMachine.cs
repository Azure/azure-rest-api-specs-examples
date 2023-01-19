using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SqlVirtualMachine;
using Azure.ResourceManager.SqlVirtualMachine.Models;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/ListBySqlVirtualMachineGroupSqlVirtualMachine.json
// this example is just showing the usage of "SqlVirtualMachines_ListBySqlVmGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlVmGroupResource created on azure
// for more information of creating SqlVmGroupResource, please refer to the document of SqlVmGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlVmGroupName = "testvm";
ResourceIdentifier sqlVmGroupResourceId = SqlVmGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlVmGroupName);
SqlVmGroupResource sqlVmGroup = client.GetSqlVmGroupResource(sqlVmGroupResourceId);

// invoke the operation and iterate over the result
await foreach (SqlVmResource item in sqlVmGroup.GetSqlVmsBySqlVmGroupAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SqlVmData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
