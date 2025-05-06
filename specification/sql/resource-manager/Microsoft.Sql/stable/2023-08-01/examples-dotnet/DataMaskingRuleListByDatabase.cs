using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/DataMaskingRuleListByDatabase.json
// this example is just showing the usage of "DataMaskingRules_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataMaskingPolicyResource created on azure
// for more information of creating DataMaskingPolicyResource, please refer to the document of DataMaskingPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-6852";
string serverName = "sqlcrudtest-2080";
string databaseName = "sqlcrudtest-331";
DataMaskingPolicyName dataMaskingPolicyName = DataMaskingPolicyName.Default;
ResourceIdentifier dataMaskingPolicyResourceId = DataMaskingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName, dataMaskingPolicyName);
DataMaskingPolicyResource dataMaskingPolicy = client.GetDataMaskingPolicyResource(dataMaskingPolicyResourceId);

// invoke the operation and iterate over the result
await foreach (DataMaskingRule item in dataMaskingPolicy.GetDataMaskingRulesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
