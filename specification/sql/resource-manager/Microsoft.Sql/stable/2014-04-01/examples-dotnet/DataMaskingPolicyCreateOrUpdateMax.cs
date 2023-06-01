using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingPolicyCreateOrUpdateMax.json
// this example is just showing the usage of "DataMaskingPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
ResourceIdentifier dataMaskingPolicyResourceId = DataMaskingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
DataMaskingPolicyResource dataMaskingPolicy = client.GetDataMaskingPolicyResource(dataMaskingPolicyResourceId);

// invoke the operation
DataMaskingPolicyData data = new DataMaskingPolicyData()
{
    DataMaskingState = DataMaskingState.Enabled,
    ExemptPrincipals = "testuser;",
};
ArmOperation<DataMaskingPolicyResource> lro = await dataMaskingPolicy.CreateOrUpdateAsync(WaitUntil.Completed, data);
DataMaskingPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataMaskingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
