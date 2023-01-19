using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SqlVirtualMachine;
using Azure.ResourceManager.SqlVirtualMachine.Models;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/CreateOrUpdateSqlVirtualMachineGroup.json
// this example is just showing the usage of "SqlVirtualMachineGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SqlVmGroupResource
SqlVmGroupCollection collection = resourceGroupResource.GetSqlVmGroups();

// invoke the operation
string sqlVmGroupName = "testvmgroup";
SqlVmGroupData data = new SqlVmGroupData(new AzureLocation("northeurope"))
{
    SqlImageOffer = "SQL2016-WS2016",
    SqlImageSku = SqlVmGroupImageSku.Enterprise,
    WindowsServerFailoverClusterDomainProfile = new WindowsServerFailoverClusterDomainProfile()
    {
        DomainFqdn = "testdomain.com",
        OrganizationalUnitPath = "OU=WSCluster,DC=testdomain,DC=com",
        ClusterBootstrapAccount = "testrpadmin",
        ClusterOperatorAccount = "testrp@testdomain.com",
        SqlServiceAccount = "sqlservice@testdomain.com",
        StorageAccountUri = new Uri("https://storgact.blob.core.windows.net/"),
        StorageAccountPrimaryKey = "<primary storage access key>",
        ClusterSubnetType = SqlVmClusterSubnetType.MultiSubnet,
    },
    Tags =
    {
    ["mytag"] = "myval",
    },
};
ArmOperation<SqlVmGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sqlVmGroupName, data);
SqlVmGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlVmGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
