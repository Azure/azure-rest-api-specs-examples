using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SqlVirtualMachine;
using Azure.ResourceManager.SqlVirtualMachine.Models;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/UpdateSqlVirtualMachine.json
// this example is just showing the usage of "SqlVirtualMachines_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlVmResource created on azure
// for more information of creating SqlVmResource, please refer to the document of SqlVmResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlVmName = "testvm";
ResourceIdentifier sqlVmResourceId = SqlVmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlVmName);
SqlVmResource sqlVm = client.GetSqlVmResource(sqlVmResourceId);

// invoke the operation
SqlVmPatch patch = new SqlVmPatch()
{
    Tags =
    {
    ["mytag"] = "myval",
    },
};
ArmOperation<SqlVmResource> lro = await sqlVm.UpdateAsync(WaitUntil.Completed, patch);
SqlVmResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlVmData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
