using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/DataMaskingRuleCreateOrUpdateText.json
// this example is just showing the usage of "DataMaskingRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string dataMaskingRuleName = "rule1";
DataMaskingRule dataMaskingRule = new DataMaskingRule
{
    SchemaName = "dbo",
    TableName = "Table_1",
    ColumnName = "test1",
    MaskingFunction = DataMaskingFunction.Text,
    PrefixSize = "1",
    SuffixSize = "0",
    ReplacementString = "asdf",
};
DataMaskingRule result = await dataMaskingPolicy.CreateOrUpdateDataMaskingRuleAsync(dataMaskingRuleName, dataMaskingRule);

Console.WriteLine($"Succeeded: {result}");
